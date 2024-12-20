package main

import (
	"dynamic-notification-system/config"
	"dynamic-notification-system/plugins"
	"fmt"
	"log"

	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" // For routing
	"github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron
var db *sql.DB
var notifiers []config.Notifier // Declare global notifiers

func main() {

	var err error

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Load plugins based on configuration
	notifiers, err = plugins.LoadPlugins(cfg.Channels)
	if err != nil {
		log.Fatalf("Error loading plugins: %v", err)
	}

	// Construct DB connection string
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err = sql.Open("mysql", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	cronInstance = cron.New()
	loadJobs(cronInstance)
	r := mux.NewRouter()
	r.HandleFunc("/jobs", handlePostJob).Methods("POST")
	r.HandleFunc("/jobs", handleGetJobs).Methods("GET") // Add a GET endpoint to return current jobs

	fmt.Println("Starting scheduled jobs...")
	go cronInstance.Start()

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func loadJobsFromDB(db *sql.DB) ([]config.ScheduledJob, error) {
	rows, err := db.Query("SELECT id, name, notification_type, recipient, message, schedule_expression FROM scheduled_jobs")
	if err != nil {
		return nil, fmt.Errorf("querying jobs: %w", err)
	}
	defer rows.Close()

	var jobs []config.ScheduledJob
	for rows.Next() {
		var job config.ScheduledJob
		err := rows.Scan(&job.ID, &job.Name, &job.NotificationType, &job.Recipient, &job.Message, &job.ScheduleExpression)
		if err != nil {
			return nil, fmt.Errorf("scanning job: %w", err)
		}
		jobs = append(jobs, job)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return jobs, nil
}

func loadJobs(c *cron.Cron) {
	// Load from DB
	log.Printf("=== started Loading from DB ===")
	dbJobs, err := loadJobsFromDB(db)
	if err != nil {
		log.Println("Error loading from DB:", err)
	}
	for _, job := range dbJobs {
		addCronJob(c, job, notifiers)
	}
	log.Printf("=== End Loading from DB ===")

}

func validateJob(job *config.ScheduledJob) error {
	if job.Name == "" {
		return fmt.Errorf("job name is required")
	}
	if job.ScheduleExpression == "" {
		return fmt.Errorf("schedule expression is required")
	}
	// Add further validation as needed
	return nil
}

func handlePostJob(w http.ResponseWriter, r *http.Request) {
	var job config.ScheduledJob
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validation
	if err := validateJob(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert into the database
	result, err := db.Exec("INSERT INTO scheduled_jobs (name, notification_type, recipient, message, schedule_expression) VALUES (?, ?, ?, ?, ?)",
		job.Name, job.NotificationType, job.Recipient, job.Message, job.ScheduleExpression)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting job: %v", err), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	job.ID = int(id)

	addCronJob(cronInstance, job, notifiers)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job) // Return the created job with ID
}

func handleGetJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []config.ScheduledJob

	rows, err := db.Query("SELECT id, name, notification_type, recipient, message, schedule_expression FROM scheduled_jobs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var job config.ScheduledJob

		err := rows.Scan(&job.ID, &job.Name, &job.NotificationType, &job.Recipient, &job.Message, &job.ScheduleExpression)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jobs = append(jobs, job)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func addCronJob(c *cron.Cron, job config.ScheduledJob, notifiers []config.Notifier) {
	jobCopy := job
	_, err := c.AddFunc(job.ScheduleExpression, func() {
		// Logic
		for _, notifier := range notifiers {
			if notifier.Type() == jobCopy.NotificationType {
				fmt.Printf("Running job: %s for %s\n", jobCopy.Name, jobCopy.Recipient)
				err := notifier.Notify(&jobCopy.Message)
				if err != nil {
					log.Printf("Error sending notification via %s: %v", notifier.Name(), err)
				}
				_, err = db.Exec("UPDATE scheduled_jobs SET last_run = NOW() WHERE id = ?", jobCopy.ID)
				if err != nil {
					log.Printf("Error updating last_run: %v", err)
				}
			}
		}

	})
	if err != nil {
		log.Printf("Error adding cron job: %v", err)
	} else {
		log.Printf("Added cron job: %s", job.Name)
	}
}

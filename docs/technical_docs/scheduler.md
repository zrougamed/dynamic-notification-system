# Scheduler Module 🕒

## Overview
The Scheduler Module handles job scheduling and execution in the **Dynamic Notification System**. It is responsible for initializing the cron-based scheduler, managing scheduled jobs, and interacting with the database for job persistence.

---

## scheduler.go 🛠️

### Purpose
Manages the lifecycle of the scheduler, including initialization and shutdown.

### Key Functions 🔑

#### **Initialize**
- **Purpose**: Sets up the scheduler and loads jobs from the database.
- **Steps**:
  1. Establishes a connection to the database using the configuration settings.
  2. Initializes a cron scheduler.
  3. Loads jobs from the database and adds them to the scheduler.
  4. Starts the scheduler.
- **Example**:
  ```go
  func Initialize(dbConfig DatabaseConfig) error {
      db, err := sql.Open("postgres", getDSN(dbConfig))
      if err != nil {
          return err
      }
      cron := cron.New()
      err = loadJobs(db, cron)
      if err != nil {
          return err
      }
      cron.Start()
      return nil
  }
  ```

#### **Shutdown**
- **Purpose**: Gracefully stops the scheduler and closes the database connection.
- **Steps**:
  1. Stops the cron scheduler to prevent further job execution.
  2. Closes the database connection.
- **Example**:
  ```go
  func Shutdown() {
      cron.Stop()
      if db != nil {
          db.Close()
      }
  }
  ```

---

## job.go 📝

### Purpose
Handles job-related functionalities such as loading, adding, and managing jobs.

### Key Functions 🔑

#### **GetJobSchema**
- **Purpose**: Generates and returns a JSON schema for the `ScheduledJob` struct, aiding in API documentation or validation.
- **Example**:
  ```go
  func GetJobSchema() string {
      reflector := jsonschema.Reflector{}
      schema, _ := reflector.Reflect(&ScheduledJob{})
      return schema.String()
  }
  ```

#### **loadJobs**
- **Purpose**: Loads jobs from the database and schedules them in the cron instance.
- **Steps**:
  1. Queries the `scheduled_jobs` table for active jobs.
  2. Adds each job to the scheduler.
- **Example**:
  ```go
  func loadJobs(db *sql.DB, cron *cron.Cron) error {
      rows, err := db.Query("SELECT * FROM scheduled_jobs")
      if err != nil {
          return err
      }
      defer rows.Close()
      for rows.Next() {
          job := ScheduledJob{}
          rows.Scan(&job)
          addCronJob(cron, job)
      }
      return nil
  }
  ```

#### **addCronJob**
- **Purpose**: Adds a job to the cron scheduler based on its cron expression.
- **Example**:
  ```go
  func addCronJob(cron *cron.Cron, job ScheduledJob) {
      cron.AddFunc(job.ScheduleExpression, func() {
          Notify(job)
      })
  }
  ```

#### **HandlePostJob**
- **Purpose**: HTTP handler for creating a new scheduled job via a POST request.
- **Example**:
  ```go
  func HandlePostJob(w http.ResponseWriter, r *http.Request) {
      var job ScheduledJob
      json.NewDecoder(r.Body).Decode(&job)
      // Insert job into the database and schedule it
      addCronJob(cron, job)
  }
  ```

#### **HandleGetJobs**
- **Purpose**: HTTP handler for retrieving all scheduled jobs via a GET request.
- **Example**:
  ```go
  func HandleGetJobs(w http.ResponseWriter, r *http.Request) {
      jobs := getJobsFromDB()
      json.NewEncoder(w).Encode(jobs)
  }
  ```

---

## db.go 🗄️

### Purpose
Handles database interactions for job scheduling.

### Key Functions 🔑

#### **loadJobsFromDB**
- **Purpose**: Queries the `scheduled_jobs` table and returns all scheduled jobs.
- **Steps**:
  1. Executes a SQL query to fetch jobs.
  2. Scans the result set into `ScheduledJob` structs.
  3. Returns the jobs for further processing.
- **Example**:
  ```go
  func loadJobsFromDB(db *sql.DB) ([]ScheduledJob, error) {
      rows, err := db.Query("SELECT * FROM scheduled_jobs")
      if err != nil {
          return nil, err
      }
      defer rows.Close()
      var jobs []ScheduledJob
      for rows.Next() {
          var job ScheduledJob
          rows.Scan(&job)
          jobs = append(jobs, job)
      }
      return jobs, nil
  }
  ```

---

## Example Workflow 🔄

1. **Initialization**:
   - The scheduler connects to the database and loads jobs into the cron instance.
   - The cron scheduler starts executing jobs based on their defined schedule.

2. **Adding Jobs**:
   - A new job is added via the `POST /jobs` endpoint.
   - The job is inserted into the database and added to the scheduler.

3. **Execution**:
   - The cron scheduler triggers the job's execution at the defined time.
   - Notifications are sent via the appropriate channel.

4. **Shutdown**:
   - The scheduler stops gracefully, ensuring no running jobs are interrupted.
   - Database connections are closed.

---

This documentation provides a detailed explanation of the Scheduler Module, enabling you to understand and extend its functionality effectively. Happy scheduling! 🎉

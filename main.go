package main

import (
	"dynamic-notification-system/config"
	"dynamic-notification-system/plugins"
	"dynamic-notification-system/scheduler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var err error

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Load plugins based on configuration
	notifiers, err := plugins.LoadPlugins(cfg.Channels)
	if err != nil {
		log.Fatalf("Error loading plugins: %v", err)
	}

	// Initialize Scheduler if enabled
	if cfg.Scheduler {
		fmt.Println("Starting scheduled jobs...")
		err = scheduler.Initialize(cfg, notifiers)
		if err != nil {
			log.Fatalf("Error initializing scheduler: %v", err)
		}
		defer scheduler.Shutdown()
	} else {
		fmt.Println("Scheduler is disabled in the configuration.")
	}

	r := mux.NewRouter()
	if cfg.Scheduler {
		r.HandleFunc("/schema/job", scheduler.GetJobSchema())
		r.HandleFunc("/jobs", scheduler.HandlePostJob).Methods("POST")
		r.HandleFunc("/jobs", scheduler.HandleGetJobs).Methods("GET")
	} else {
		fmt.Println("Scheduling endpoints are disabled.")
	}

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

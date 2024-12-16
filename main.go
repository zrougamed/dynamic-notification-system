package main

import (
	"dynamic-notification-system/config"
	"dynamic-notification-system/plugins"
	"dynamic-notification-system/scheduler"
	"fmt"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Load plugins based on configuration
	notifiers, err := plugins.LoadPlugins(cfg.Channels)
	if err != nil {
		log.Fatalf("Error loading plugins: %v", err)
	}

	// Initialize the scheduler
	sched := scheduler.NewScheduler()

	// Schedule jobs for each notifier
	for _, notifier := range notifiers {
		err := sched.ScheduleJob(notifier)
		if err != nil {
			log.Printf("Error scheduling job for %s: %v", notifier.Name(), err)
		}
	}

	// Start the scheduler
	fmt.Println("Starting scheduled jobs...")
	sched.Start()
}

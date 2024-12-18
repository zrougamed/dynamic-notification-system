package scheduler

import (
	"dynamic-notification-system/config"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		scheduler: gocron.NewScheduler(time.UTC),
	}
}

func (s *Scheduler) ScheduleJob(notifier config.Notifier) error {
	_, err := s.scheduler.Every(1).Minute().Do(func() {

		message := &config.Message{
			Title:    "Server Alert 🚨",
			Text:     "Disk usage is above 90% on server 'prod-01'.",
			Tags:     []string{"warning", "server"}, // Tags as a JSON array
			Priority: 5,                             // Max priority
			Attach:   "https://example.com/logs/error.log",
			Email:    "admin@example.com",
			Actions: []interface{}{
				map[string]string{
					"action": "view",
					"label":  "View Logs",
					"url":    "https://example.com/logs/error.log",
				},
				map[string]string{
					"action": "http", // A valid action
					"label":  "Acknowledge",
					"url":    "https://example.com/acknowledge",
				},
			},
		}

		err := notifier.Notify(message)
		if err != nil {
			log.Printf("Error sending notification via %s: %v", notifier.Name(), err)
		}
	})
	return err
}

func (s *Scheduler) Start() {
	s.scheduler.StartBlocking()
}

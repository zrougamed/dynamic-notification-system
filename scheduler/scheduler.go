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
		err := notifier.Notify("Scheduled message!")
		if err != nil {
			log.Printf("Error sending notification via %s: %v", notifier.Name(), err)
		}
	})
	return err
}

func (s *Scheduler) Start() {
	s.scheduler.StartBlocking()
}

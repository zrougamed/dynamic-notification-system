package notifier

import (
	"dynamic-notification-system/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var notifiers []config.Notifier

// SetNotifiers initializes the notifiers
func SetNotifiers(n []config.Notifier) {
	notifiers = n
}

func HandlePostJob(w http.ResponseWriter, r *http.Request) {
	var job config.InstantJob

	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateJob(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jobCopy := job
	for _, notifier := range notifiers {
		if notifier.Type() == jobCopy.NotificationType {
			fmt.Printf("Running job: %s \n", jobCopy.Recipient)
			err := notifier.Notify(&jobCopy.Message)
			if err != nil {
				log.Printf("Error sending notification via %s: %v", notifier.Name(), err)
			}
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

func validateJob(job *config.InstantJob) error { // add instant notification
	if job.NotificationType == "" {
		return fmt.Errorf("NotificationType is required")
	}
	return nil
}

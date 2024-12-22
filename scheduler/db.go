package scheduler

import (
	"database/sql"
	"dynamic-notification-system/config"
	"fmt"
)

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

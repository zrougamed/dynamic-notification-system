package scheduler

import (
	"dynamic-notification-system/config"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
)

var cronInstance *cron.Cron
var db *sql.DB
var notifiers []config.Notifier

// Initialize sets up the database and cron instance.
func Initialize(cfg *config.Config, loadedNotifiers []config.Notifier) error {
	var err error
	notifiers = loadedNotifiers

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
		return fmt.Errorf("failed to connect to DB: %w", err)
	}

	cronInstance = cron.New()
	loadJobs(cronInstance)
	go cronInstance.Start()
	return nil
}

// Shutdown gracefully stops the cron instance and closes the database.
func Shutdown() {
	cronInstance.Stop()
	if db != nil {
		db.Close()
	}
}

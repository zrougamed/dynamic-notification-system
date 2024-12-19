CREATE TABLE IF NOT EXISTS scheduled_jobs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    notification_type VARCHAR(255) NOT NULL,
    recipient VARCHAR(255) NOT NULL,
    message TEXT,
    schedule_expression VARCHAR(255) NOT NULL,
    last_run DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert dummy data (optional)
INSERT INTO scheduled_jobs (name, notification_type, recipient, message, schedule_expression) VALUES
('Daily Report', 'email', 'report@example.com', 'Daily report email', '0 0 * * *'),
('Hourly Update', 'sms', '+15551234567', 'Hourly update SMS', '0 * * * *');
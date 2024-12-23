# Developer Guide

## Architecture Overview 🔍

### Main Application

The core of the application, responsible for:

  - Loading and validating the configuration (`config.yaml`).
  - Dynamically loading plugins for different notification channels.
  - Initializing and managing the scheduler for job execution.
  - Setting up the HTTP server to handle RESTful APIs for managing jobs.

### Plugins
Plugins are modular, dynamically loaded components that extend the application's notification capabilities:

  - **Dynamic Loading**: Add new plugins without restarting the application.
  - **Interface Compliance**: Each plugin must implement the `Notifier` interface:

  ```go
  type Notifier interface {
      Name() string
      Type() string
      Notify(message *Message) error
  }
  ```

  - Examples of plugins include Email, Slack, SMS, and Webhook notifiers.

### Scheduler
The scheduler is built using a cron-based mechanism to manage job execution:

  - **Job Management**: Handles adding, removing, and executing jobs based on cron expressions.
  - **Persistence**: Stores job details in a database for reliable scheduling.
  - **Concurrency**: Supports concurrent job execution while maintaining thread safety.

---

## Plugin Development 🛠️

### Steps to Develop a New Plugin

1. **Implement the Interface**:

      - Create a new Go file for your plugin and implement the `Notifier` interface.
      - Example:

     ```go
     type MyPlugin struct {}

     func (p *MyPlugin) Name() string {
         return "MyPlugin"
     }

     func (p *MyPlugin) Type() string {
         return "Custom"
     }

     func (p *MyPlugin) Notify(message *Message) error {
         // Your notification logic
         return nil
     }
     ```

2. **Compile the Plugin**:

     - Use Go's plugin build mode to compile your plugin into a shared object file:
     ```bash
     go build -buildmode=plugin -o plugins/my_plugin.so my_plugin.go
     ```

3. **Test the Plugin**:

    - Ensure your plugin works as expected by integrating it into the application.
    - Example usage:
     ```go
     plugin, err := plugin.Open("plugins/my_plugin.so")
     if err != nil {
         log.Fatal(err)
     }
     // Further plugin initialization...
     ```

4. **Integrate with the Application**:

     - Update the `config.yaml` to include your new plugin.

---

## Database Schema 📊

### Scheduled Jobs Table
The `scheduled_jobs` table stores information about scheduled notifications:

  - **Columns**:
    - `id`: Unique identifier for the job.
    - `name`: Name of the scheduled job.
    - `notification_type`: Type of notification (e.g., email, Slack).
    - `recipient`: Target recipient (e.g., email address, phone number).
    - `message`: JSON-encoded message to be sent.
    - `schedule_expression`: Cron expression defining the job schedule.

---

## Deployment 🚀

### Using Docker
1. **Build Docker Image**:
   ```bash
   docker build -t dynamic-notification-system .
   ```
2. **Run Container**:
   ```bash
   docker run -d -p 8080:8080 dynamic-notification-system
   ```

### Without Docker
1. Compile the application:
   ```bash
   make all
   cd build
   ```
2. Run the application:
   ```bash
   ./notification-system
   ```

---

## API Endpoints 🌐

### Job Management

  - **POST /jobs**:
    - Adds a new job.
    - Example Request:
      ```json
      {
          "name": "Daily Report",
          "notification_type": "email",
          "recipient": "user@example.com",
          "message": {
              "title": "Daily Report",
              "message": "Your daily report is ready."
          },
          "schedule_expression": "0 9 * * *"
      }
      ```

  - **GET /jobs**:
    - Retrieves all scheduled jobs.

  - **GET /schema/job**:
    - Retrieves the job schema for easier integration and validation.

---

Thank you for contributing to and developing the Dynamic Notification System! 🎉

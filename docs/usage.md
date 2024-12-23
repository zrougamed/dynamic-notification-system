# Usage 📘

Learn how to make the most of the **Dynamic Notification System** by adding new notification channels and scheduling notifications.

---

## Adding New Notification Channels 🔔

The system supports a plugin-based architecture for notification channels. Here’s how to add a new channel:

1. **Create a Go Plugin**:

    - Develop a Go plugin that implements the `Notifier` interface:
    ```go
    type Notifier interface {
        Name() string
        Type() string
        Notify(message *Message) error
    }
    ```

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
        fmt.Printf("Sending message: %s", message.Text)
        return nil
    }
    ```

2. **Compile the Plugin**:

    - Use the following command to compile your Go plugin:
      ```bash
      go build -buildmode=plugin -o plugins/my_plugin.so my_plugin.go
      ```

3. **Add the Plugin to the Directory**:

    - Place the compiled `.so` file into the `plugins/` directory.

4. **Update `config.yaml`**:

    - Add the new channel configuration to `config.yaml`. Example:
      ```yaml
      channels:
        my_plugin:
          enabled: true
          custom_field: "value"
      ```

5. **Restart the Application**:

    - Restart the application to load the new plugin:
      ```bash
      ./notification-system
      ```

---

## Scheduling Notifications 🗓️

The scheduler enables you to define and automate notification jobs.

### Using the `/jobs` API Endpoint
You can define and schedule jobs directly via the HTTP API.

1. **POST /jobs**:

    - Use the `/jobs` endpoint to create a new job. Example request:
      ```bash
      curl -X POST http://localhost:8080/jobs      -H "Content-Type: application/json"      -d '{
          "name": "Daily Report",
          "notification_type": "email",
          "recipient": "user@example.com",
          "message": {
              "title": "Daily Report",
              "message": "Your daily report is ready."
          },
          "schedule_expression": "0 9 * * *"
      }'
      ```

2. **Verify the Job**:

    - Check the response to confirm the job is created:
      ```json
      {
          "id": 1,
          "name": "Daily Report",
          "notification_type": "email",
          "recipient": "user@example.com",
          "schedule_expression": "0 9 * * *"
      }
      ```

3. **Scheduler Execution**:

    - The scheduler will execute the job at the defined time based on the cron expression.

---

## Advanced Usage ⚙️

### Editing Jobs:

  - Modify job details directly via the database or through future API endpoints:
  ```sql
  UPDATE scheduled_jobs
  SET schedule_expression = '0 10 * * *'
  WHERE id = 1;
  ```

### Deleting Jobs:

  - Remove a job from the schedule:
  ```sql
  DELETE FROM scheduled_jobs WHERE id = 1;
  ```

---

## Examples ✨

### Example: Adding a Slack Notification Job
1. **Add a Slack plugin** using the steps above.

2. **Use the `/jobs` API endpoint** to add a Slack notification job:
  ```bash
  curl -X POST http://localhost:8080/jobs \
  -H "Content-Type: application/json" \
  -d '{
      "name": "Team Standup Reminder",
      "notification_type": "slack",
      "recipient": "#general",
      "message": {
          "title": "Standup Reminder",
          "message": "Daily standup in 10 minutes!"
      },
      "schedule_expression": "30 8 * * 1-5"
  }'
  ```

3. **Verify execution**:
    - The message will be sent to the `#general` Slack channel at **8:30 AM**, Monday to Friday.

---

Enjoy using the Dynamic Notification System to streamline your notifications! 🚀

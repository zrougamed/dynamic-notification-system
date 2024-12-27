
# Notifier Package

The **Notifier** package is responsible for handling dynamic notifications within the system. It supports instant notifications via an HTTP endpoint and integrates seamlessly with the existing system configuration.

## Key Features

- **Instant Notifications**: Handles and processes notification jobs dynamically.
- **Flexible Integration**: Uses configurable notifiers to send notifications based on the `NotificationType`.
- **Validation**: Ensures proper validation of job requests before execution.

## Components

### 1. SetNotifiers
Initializes the available notifiers.

```go
func SetNotifiers(n []config.Notifier) {
    notifiers = n
}
```

### 2. HandlePostJob
Handles HTTP POST requests for instant notifications.

```go
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
    for _, notifier := range notifiers {
        if notifier.Type() == job.NotificationType {
            err := notifier.Notify(&job.Message)
            if err != nil {
                log.Printf("Error with %s: %v", notifier.Name(), err)
            }
        }
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(job)
}
```

### 3. Validation
Validates the `InstantJob` struct before processing.

```go
func validateJob(job *config.InstantJob) error {
    if job.NotificationType == "" {
        return fmt.Errorf("NotificationType is required")
    }
    return nil
}
```

## Usage

1. Configure the notifiers using `SetNotifiers` with a list of notifier implementations.
2. Use the `/notify` endpoint to send instant notifications with the required payload.

## Example Payload

```json
{
    "notificationType": "ntfy",
    "recipient": "user@example.com",
    "message": {
        "title": "Disk Usage 🚨",
        "message": "Disk usage is above 90% on server prod-01.",
        "tags": ["warning", "server"],
        "priority": 5,
        "attach": "https://example.com/logs/error.log",
        "email": "admin@example.com",
        "actions": [
        {
            "action": "view",
            "label": "View Logs",
            "url": "https://example.com/logs/error.log"
        },
        {
            "action": "http",
            "label": "Acknowledge",
            "url": "https://example.com/acknowledge"
        }
        ]
    }
}

```

---

For more details on configuration, refer to [config](config.md).

This documentation provides a detailed explanation of the Instant notification Module, enabling you to understand and extend its functionality effectively. Happy notifications! 🎉


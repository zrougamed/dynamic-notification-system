# main.go 🛠️

## Purpose
The `main.go` file serves as the entry point of the **Dynamic Notification System** application. It orchestrates the initialization of configurations, plugins, the scheduler, and the HTTP server.

---

## Key Functions 🔑

### 1. **Configuration Loading**
- **Purpose**: Loads the application's configuration settings from the `config.yaml` file.
- **Implementation**:
  - Function: `config.LoadConfig`
  - Steps:
    1. Reads the `config.yaml` file from the root directory.
    2. Validates and parses the configuration into a structured `Config` object.
  - Example:
    ```go
    config, err := config.LoadConfig("config.yaml")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
    ```

### 2. **Plugin Loading**
- **Purpose**: Dynamically loads plugins that define notification channels.
- **Implementation**:
  - Function: `plugins.LoadPlugins`
  - Steps:
    1. Scans the `plugins/` directory for `.so` files.
    2. Loads and initializes each plugin dynamically.
  - Example:
    ```go
    err := plugins.LoadPlugins("plugins/")
    if err != nil {
        log.Fatalf("Failed to load plugins: %v", err)
    }
    ```

### 3. **Scheduler Initialization**
- **Purpose**: Sets up the job scheduler, responsible for executing tasks based on cron expressions.
- **Implementation**:
  - Function: `scheduler.Initialize`
  - Steps:
    1. Establishes a database connection.
    2. Loads jobs from the database into the scheduler.
    3. Starts the cron scheduler.
  - Example:
    ```go
    err := scheduler.Initialize(config.Database)
    if err != nil {
        log.Fatalf("Failed to initialize scheduler: %v", err)
    }
    ```

### 4. **HTTP Server Setup**
- **Purpose**: Starts the HTTP server to handle API requests for managing notifications and jobs.
- **Implementation**:
  - Library: Gorilla Mux
  - Steps:
    1. Creates a new router.
    2. Defines routes for job management endpoints (`/jobs`).
    3. Starts the server on the configured port (default: 8080).
  - Example:
    ```go
    r := mux.NewRouter()
    r.HandleFunc("/jobs", HandlePostJob).Methods("POST")
    r.HandleFunc("/jobs", HandleGetJobs).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", r))
    ```

---

## Example Flow 🔄

1. **Application Start**:
   - Reads `config.yaml`.
   - Loads plugins from the `plugins/` directory.
   - Initializes the scheduler and loads jobs.
   - Starts the HTTP server.

2. **Incoming API Requests**:
   - Routes are handled via the Gorilla Mux router.
   - Jobs are added, retrieved, or managed based on the API calls.

---

## Error Handling ⚠️
- Logs detailed error messages and exits gracefully if critical components fail to initialize.
- Example:
  ```go
  if err := scheduler.Initialize(config.Database); err != nil {
      log.Fatalf("Scheduler initialization failed: %v", err)
  }
  ```

---

This documentation provides a high-level overview and examples for understanding and extending the functionality of the `main.go` file. Happy coding! 🚀

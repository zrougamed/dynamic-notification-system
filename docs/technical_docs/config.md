# Config Module 📋

## config.go 🛠️

### Purpose
The `config.go` file defines the application's configuration and provides utilities to load it from a YAML file. This module is essential for managing dynamic settings across various components.

---

## Key Components 🔑

### 1. **Message Struct**
- **Purpose**: Represents notification messages used across the system.
- **Key Features**:
  - Implements `sql.Scanner` to parse JSON from the database into a `Message` struct.
  - Implements `driver.Valuer` to convert the `Message` struct into JSON for database storage.
- **Example**:
  ```go
  type Message struct {
      Title string `json:"title"`
      Text  string `json:"text"`
  }

  func (m *Message) Scan(value interface{}) error {
      return json.Unmarshal(value.([]byte), m)
  }

  func (m Message) Value() (driver.Value, error) {
      return json.Marshal(m)
  }
  ```

### 2. **ScheduledJob Struct**
- **Purpose**: Defines the fields required for scheduling notification jobs.
- **Key Fields**:
  - `ID`: Unique identifier for the job.
  - `Name`: Name of the job.
  - `NotificationType`: Type of notification (e.g., email, Slack).
  - `Recipient`: Target recipient.
  - `Message`: Message content in `Message` struct format.
  - `ScheduleExpression`: Cron expression for job timing.
- **Example**:
  ```go
  type ScheduledJob struct {
      ID                 int       `json:"id"`
      Name               string    `json:"name"`
      NotificationType   string    `json:"notification_type"`
      Recipient          string    `json:"recipient"`
      Message            Message   `json:"message"`
      ScheduleExpression string    `json:"schedule_expression"`
  }
  ```

### 3. **Notifier Interface**
- **Purpose**: Defines the methods that custom notifiers must implement.
- **Methods**:
  - `Name() string`: Returns the name of the notifier.
  - `Type() string`: Returns the type of the notifier.
  - `Notify(message *Message) error`: Sends a notification based on the provided message.
- **Example Implementation**:
  ```go
  type Notifier interface {
      Name() string
      Type() string
      Notify(message *Message) error
  }
  ```

### 4. **LoadConfig Function**
- **Purpose**: Reads and parses the configuration file into a structured `Config` object.
- **Steps**:
  1. Opens the `config.yaml` file.
  2. Parses its content into a `Config` struct.
  3. Returns the structured configuration for use across the application.
- **Example**:
  ```go
  func LoadConfig(path string) (*Config, error) {
      data, err := ioutil.ReadFile(path)
      if err != nil {
          return nil, err
      }
      var config Config
      err = yaml.Unmarshal(data, &config)
      if err != nil {
          return nil, err
      }
      return &config, nil
  }
  ```

---

## Config Struct 📦

### Purpose
Encapsulates the entire configuration for the application, including database settings, notification channels, and scheduler options.

### Key Fields
- **Database Config**:
  - `Host`, `Port`, `User`, `Password`, `Name`
- **Channels Config**:
  - Email, Slack, SMS, Webhook configurations
- **Scheduler Flag**:
  - Enables or disables the job scheduler.

### Example
```go
type Config struct {
    Database  DatabaseConfig    `yaml:"database"`
    Channels  map[string]ChannelConfig `yaml:"channels"`
    Scheduler bool              `yaml:"scheduler"`
}

type DatabaseConfig struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    Name     string `yaml:"name"`
}

type ChannelConfig struct {
    Enabled    bool   `yaml:"enabled"`
    Host       string `yaml:"host"`
    Port       int    `yaml:"port"`
    Username   string `yaml:"username"`
    Password   string `yaml:"password"`
    WebhookURL string `yaml:"webhook_url"`
}
```

---

## Example Flow 🔄

1. **Load Configuration**:
   - The application reads `config.yaml` on startup.
   - Parses the YAML file into a structured `Config` object.
2. **Use Configuration**:
   - Components access configuration values directly from the `Config` object.
   - Example:
     ```go
     dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
      cfg.Database.User,
      cfg.Database.Password,
      cfg.Database.Host,
      cfg.Database.Port,
      cfg.Database.Name,
    )
     ```

---

This documentation provides a detailed explanation of `config.go`, highlighting its essential role in managing the application’s dynamic settings. 🛠️ Happy configuring!

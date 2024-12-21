# 📤 Dynamic Notification System Documentation

## 📝 Overview
The **Dynamic Notification System** is a platform designed for managing ⏰ jobs and sending 📧 via multiple 📡. It supports:
- 📥 dynamic 🔌 for various 📬 methods.
- Managing ⏰ jobs through a ⏱️-based scheduler.
- 🌐 APIs for creating and managing jobs.

---

## ⭐ Features
- **🔄 Dynamic Plugin System**: Load 📬 🔌 dynamically based on the ⚙️.
- **🗄️ Database Integration**: Store ⏰ jobs and their execution 📜 in a 🐬 database.
- **⏱️ Scheduler**: Manage job ⏰ and ensure timely execution.
- **🌐 API**:
  - ✏️ new jobs.
  - 📄 existing jobs.
- **⚠️ Error Handling**: Gracefully handle 🚨 and 🔗.

---

## 🏗️ Architecture
### 🧩 Components
1. **⚙️ Configuration Loader**:
   - 🛠️ settings from a 🗂️ YAML ⚙️ file.
   - Configures 🐬, 🔌, and 🖥️ behavior.

2. **🗄️ Database**:
   - 🛠️ job metadata, including ⏰ expressions, 📬, and execution 📜.

3. **🔌 Plugins**:
   - Dynamically loaded 📬 🔌 to send 📨 through different 📡 (e.g., 📧, 📩, 💬).

4. **⏱️ Scheduler**:
   - Based on the `cron` 🛠️.
   - Executes ⏰ at specified 🕒.

5. **🌐 API**:
   - Built with `Gorilla Mux` for 🛣️.
   - Provides 🔗 for job management.

---

## 🛠️ Setup and Installation

### 🧾 Prerequisites
- 🐹 Go (latest version)
- 🐬 MySQL database ( docker compose includes one )
- 🗂️ YAML ⚙️ file

### 📦 Installation Steps
1. 🌀 the 📂:
   ```bash
   git clone https://github.com/your-repo/dynamic-notification-system.git
   cd dynamic-notification-system
   ```
2. 🛠️ the 🛠️:
   ```bash
   make all 
   ```
3. ⚙️ the 🖥️:
   - Create a `config.yaml` 🗂️:
     ```yaml
     database:
        host: localhost
        port: 3306
        user: root
        password: password
        name: notifications
     channels:
        sms:
            enabled: false
            provider_api: "https://sms-provider.com/api"
            api_key: "your-sms-api-key"
            phone_number: "recipient-phone-number"

        signal:
            enabled: false
            api_url: "https://signal-server.com"
            phone_number: "recipient-phone-number"

        rocketchat:
            enabled: false
            webhook_url: "https://chat.example.com/hooks/your-webhook-url"
     ```
4. ▶️ the 🖥️:
   ```bash
   ./dynamic-notification-system
   ```

---

## 🌐 API Endpoints
### **📤 POST /jobs**
- **📄 Description**: ✏️ a new ⏰ job.
- **📝 Request Body**:
  ```json
  {
    "name": "Job Name",
    "notification_type": "email",
    "recipient": "example@example.com",
    "message": {
            "title": "Server Alert 🚨",
            "tags": [
                "warning",
                "server"
            ],
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
    },
    "schedule_expression": "* * * * *"
  }
  ```
- **📬 Response**:
  ```json
  {
    "id": 1,
    "name": "Job Name",
    "notification_type": "email",
    "recipient": "example@example.com",
    "message": {
            "title": "Server Alert 🚨",
            "tags": [
                "warning",
                "server"
            ],
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
    },
    "schedule_expression": "* * * * *"
  }
  ```

### **📄 GET /jobs**
- **📄 Description**: 📄 all ⏰ jobs.
- **📬 Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Job Name",
      "notification_type": "email",
      "recipient": "example@example.com",
      "message": {
            "title": "Server Alert 🚨",
            "tags": [
                "warning",
                "server"
            ],
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
      },
      "schedule_expression": "* * * * *"
    }
  ]
  ```

---

## ⚙️ Configuration
### 🗂️ YAML Structure
```yaml
database:
  host: <db_host>
  port: <db_port>
  user: <db_user>
  password: <db_password>
  name: <db_name>

channels:
  sms:
    enabled: false
    provider_api: "https://sms-provider.com/api"
    api_key: "your-sms-api-key"
    phone_number: "recipient-phone-number"

  signal:
    enabled: false
    api_url: "https://signal-server.com"
    phone_number: "recipient-phone-number"

  rocketchat:
    enabled: false
    webhook_url: "https://chat.example.com/hooks/your-webhook-url"
```

---

## 🛠️ Development
### 🔑 Libraries Used
- **`github.com/gorilla/mux`**: For 🛣️.
- **`github.com/robfig/cron/v3`**: For ⏰ job ⏱️.
- **`github.com/go-sql-driver/mysql`**: 🐬 database 🔌.

### Adding a New 🔌
1. Create a new 🔌 in the `plugins` 📂.
2. 🛠️ the `Notifier` interface:
   ```go
   type Notifier interface {
       Notify(message *string) error
       Type() string
       Name() string
   }
   ```
3. Update the `config.yaml` 🛠️ template to include your 🔌.

---
## 🚀 Future Enhancements
- Support for additional 📬 📡 (e.g., WhatsApp, Telegram).
- Role-based 🛡️ control (RBAC) for 🌐 🔗.
- Enhanced ⏰ 📈 and 🔄 🛠️.

---

## 📜 License
[MIT License](https://github.com/zrougamed/dynamic-notification-system/blob/main/LICENSE)

---

## 👥 Contributors
- [Mohamed Zrouga](https://github.com/zrougamed)

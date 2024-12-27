# 📤 Dynamic Notification System Documentation

## 📝 Overview

The **Dynamic Notification System** is a platform designed for managing ⏰ jobs and sending 📧 via multiple 📡. It supports:

- 📥 Dynamic 🔌 for various 📬 methods.
- Managing ⏰ jobs through a ⏱️-based scheduler.
- 🌐 APIs for creating and managing jobs.

---

## ⭐ Features

- **🔄 Multi-platform Support**: Send notifications via 📧, Slack, SMS, and Webhooks.
- **🔌 Dynamic Plugin System**: Load 📬 🔌 dynamically based on the ⚙️.
- **🗄️ Database Integration**: Store ⏰ jobs and their execution 📜 in a 🐬 database.
- **⏱️ Scheduler**: Manage job ⏰ and ensure timely execution.
- **🌐 API**:
    - ✏️ Create new jobs.
    - 📄 Retrieve existing jobs.
- **⚠️ Error Handling**: Gracefully handle 🚨 and 🔗 issues.

---

## 🏗️ Architecture
### 🧩 Components
1. **⚙️ Configuration Loader**:

    - 🛠️ Reads settings from a 🗂️ YAML ⚙️ file.
    - Configures 🐬, 🔌, and 🖥️ behavior.

2. **🗄️ Database**:

    - 🛠️ Stores job metadata, including ⏰ expressions, 📬 details, and execution 📜.

3. **🔌 Plugins**:

    - Dynamically loaded 📬 🔌 to send 📨 through different 📡 (e.g., 📧, 📩, 💬).

4. **⏱️ Scheduler**:

    - Based on the `cron` 🛠️.
    - Executes ⏰ jobs at specified 🕒.

5. **🌐 API**:

    - Built with `Gorilla Mux` for 🛣️.
    - Provides 🔗 for job management.

---

## 🛠️ Setup and Installation

### 🧾 Prerequisites
- 🐹 Go (version 1.23+)
- 🐬 MySQL database (docker compose includes one)
- 🗂️ YAML ⚙️ file

### 📦 Installation Steps
1. Clone the repository:
```bash
git clone https://github.com/zrougamed/dynamic-notification-system.git
cd dynamic-notification-system
```
2. Build the application:
```bash
make all
```
3. Configure the application:
   - Create a `config.yaml` file:

```yaml
scheduler: true
database:
    host: localhost
    port: 3306
    user: root
    password: password
    name: notifications
channels:
    email:
        enabled: true
        smtp_server: "smtp.example.com"
        smtp_port: 587
        username: "your-email@example.com"
        password: "your-password"
    sms:
        enabled: false
        provider_api: "https://sms-provider.com/api"
        api_key: "your-sms-api-key"
```
4. Run the application:
```bash
./dynamic-notification-system
```

---

## 📚 Table of Contents
1. [Getting Started](getting_started.md)
2. [Usage](usage.md)
3. [Developer Guide](developer_guide.md)
4. [Main Module](technical_docs/main.md)
5. [Scheduler Module](technical_docs/scheduler.md)
6. [Notifier Module](technical_docs/notifier.md)
7. [Config Module](technical_docs/config.md)
8. [Contributing](contributing.md)

---

## 🌐 API Endpoints

### **📤 POST /jobs**

- **📄 Description**: Create a new scheduled job.
- **📝 Request Body**:
```json
{
"name": "Job Name",
"notification_type": "email",
"recipient": "example@example.com",
"message": {
    "title": "Server Alert 🚨",
    "text": "Disk space is low on server",
    "priority": "high"
},
"schedule_expression": "0 9 * * *"
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
    "text": "Disk space is low on server",
    "priority": "high"
},
"schedule_expression": "0 9 * * *"
}
```

### **📄 GET /jobs**
- **📄 Description**: Retrieve all scheduled jobs.
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
    "text": "Disk space is low on server",
    "priority": "high"
    },
    "schedule_expression": "0 9 * * *"
}
]
```

---

## 🛠️ Development

### 🔑 Libraries Used
- **`github.com/gorilla/mux`**: For API routing.
- **`github.com/robfig/cron/v3`**: For job scheduling.
- **`github.com/go-sql-driver/mysql`**: For MySQL database integration.

---

## 🚀 Future Enhancements
- Support for additional channels (e.g., WhatsApp, Telegram).
- Role-based access control (RBAC) for APIs.
- Enhanced logging and monitoring tools.

---

## 📜 License
[MIT License](https://github.com/zrougamed/dynamic-notification-system/blob/main/LICENSE)

---

## 👥 Contributors
- [Mohamed Zrouga](https://github.com/zrougamed)


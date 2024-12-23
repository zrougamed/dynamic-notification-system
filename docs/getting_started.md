# Getting Started 🚀

Welcome to the **Dynamic Notification System**! This guide will help you set up and run the application quickly.

---

## Prerequisites 🛠️
Ensure the following tools are installed on your system:

  1. **Go (version 1.23+)**:
      - Download and install from [golang.org](https://golang.org/dl/).
      - Verify installation:
```bash
go version
```
  2. **Docker** (optional for containerized deployment):
      - Download and install from [docker.com](https://www.docker.com/products/docker-desktop).
      - Verify installation:
```bash
docker --version
```
  3. **MySQL**:
      - Install MySQL for your platform or use the docker-compose shipped with the code.
      - Load the schema from `db` folder

---

## Installation ⚙️

Follow these steps to set up the application:

1. **Clone the Repository**:
    - Clone the source code from GitHub:
```bash
git clone https://github.com/zrougamed/dynamic-notification-system.git
cd dynamic-notification-system
```

2. **Build the Application**:
    - Compile the source code into a binary:
```bash
make all
```

---

## Configuration 📝

The application relies on a `config.yaml` file for its configuration. Here's how to set it up:

1. **Edit `config.yaml`**:
    - Modify the file to include your credentials and preferences. Example:
```yaml
database:
  host: "localhost"
  port: 3306
  user: "your-db-user"
  password: "your-db-password"
  name: "dynamic_notification_system"
scheduler: true
channels:
  email:
    enabled: true
    smtp_server: "smtp.example.com"
    smtp_port: 587
    username: "your-email@example.com"
    password: "your-email-password"
```

2. **Verify Configuration**:
    - Ensure the `config.yaml` file is in the root directory of the application:
```bash
ls | grep config.yaml
```

---

## Running the Application 🏃

### Option 1: Direct Execution
Run the application directly using the compiled binary:
```bash
./notification-system
```
- The application will start on the default port (8080).
- Access the application via:
  - API: `http://localhost:8080`

### Option 2: Using Docker
For a containerized deployment:

- **Build the Docker Image**:
```bash
docker build -t dynamic-notification-system .
```
- **Run the Container**:
```bash
docker run -d -p 8080:8080 dynamic-notification-system
```
- **Verify the Container**:
    - List running containers: ```docker ps```
    - Access the application at : `http://localhost:8080`

---

## Next Steps ➡️
  After setting up the application:

  - [Configure Plugins](developer_guide.md#plugin-development) for custom notification channels.
  - Start scheduling jobs using the [API Endpoints](developer_guide.md#api-endpoints).

Enjoy using the Dynamic Notification System! 🎉

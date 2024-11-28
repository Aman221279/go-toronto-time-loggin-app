# go-toronto-time-loggin-app

# Application Setup and Running Instructions

## 1. Install Required Software

### Install Go
Download and install Go from the official website: [golang.org](https://golang.org).

### Install MySQL
Download and install MySQL from the official website: [MySQL Downloads](https://dev.mysql.com/downloads/).

### Start the MySQL Server
Ensure that your MySQL server is running.

## 2. Setting Up the Application

### Initialize Go Environment
1. Open a terminal and navigate to the application folder.
2. Run the following command to initialize the Go environment:
   ```bash
   go mod tidy
   ```

### Modify Database Configuration
1. Open the `main.go` file.
2. Update the database name in the following line to match your MySQL database:
   ```go
   db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/database")
   ```

### Run the Application
Run the application using the command:
```bash
go run main.go
```

The application will start a server at `http://localhost:8080`.

### Test the Endpoints
You can test both endpoints using the following `curl` commands:

- Current Time Endpoint:
   ```bash
   curl http://localhost:8080/currentTime
   ```
- List Times Endpoint:
   ```bash
   curl http://localhost:8080/listTimes
   ```

Logs of the API requests will be saved in the `api.log` file.

## 3. MySQL Database Setup

### Create Database and Table
Run the following SQL script to create the required database and table:

```sql
create database assignment;

use assignment;

CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
```

## 4. Running the Application Using Docker

### Install Docker
Install Docker from the official website: [Docker's Official Site](https://www.docker.com).

### Install Docker Compose
Install Docker Compose according to the instructions provided on the Docker website.

### Update Database Configuration in Docker Compose
1. Open the `docker-compose.yml` file.
2. Update the database details according to your MySQL configuration.

### Build and Start Services
Run the following command to build and start the services:
```bash
docker-compose up --build
```

### Test the Endpoints
Once the application is running, you can test both endpoints using the following `curl` commands:

- Current Time Endpoint:
   ```bash
   curl http://localhost:8080/currentTime
   ```
- List Times Endpoint:
   ```bash
   curl http://localhost:8080/listTimes
   ```

### Stop and Remove Services
To stop and remove the Docker containers, run:
```bash
docker-compose down
```

# Go Real-Time Chat Project

This is a simple 1:1 real-time chat application built with Go for the backend. This project is designed to be a comprehensive portfolio piece, covering all stages from design, backend implementation, testing, to containerization.

## Features

* **User Authentication:** Secure registration and login system using username and password. Passwords are securely hashed using **bcrypt**.
* **User Search:** Search for other users by username to start a new conversation.
* **1:1 Real-Time Messaging:** Instant message delivery between two users using **WebSockets**.
* **Message History:** All messages are persisted in a **PostgreSQL** database, allowing users to view their conversation history.
* **Structured API:** Clear and protected API endpoints using **JSON Web Tokens (JWT)**.

## Tech Stack

### Backend

* **Language:** [Go](https://golang.org/)
* **Web Framework:** [Gin](https://gin-gonic.com/)
* **Database:** [PostgreSQL](https://www.postgresql.org/)
* **ORM:** [GORM](https://gorm.io/)
* **Real-time:** [Gorilla WebSocket](http://www.gorillatoolkit.org/pkg/websocket)
* **Environment Management:** [godotenv](https://github.com/joho/godotenv)
* **Security:** [Bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt), [JWT](https://github.com/golang-jwt/jwt)

### Deployment

* **Containerization:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

## How to Run the Project

There are two main ways to run this project: using **Docker** (the recommended and easiest way) or **manually** in your local environment.

### 1. Running with Docker (Recommended)

This method will automatically build the Go application and run the PostgreSQL database in isolated containers. It's the easiest way to get started.

**Prerequisites:**

* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

**Steps:**

1.  **Clone This Repository**
    ```bash
    git clone https://github.com/lskeey/websocket-chat.git
    cd websocket-chat
    ```

2.  **Configure Environment**
    Copy the `.env.example` file to `.env`. The default values are already configured to work with Docker Compose, so you don't need to change anything.
    ```bash
    cp .env.example .env
    ```

3.  **Build and Run with Docker Compose**
    Run the following command from the project's root directory. This command will build the Docker image for the Go application and start all services.
    ```bash
    docker-compose up --build
    ```

4.  **Access the Application**
    The backend API is now running and accessible at `http://localhost:8080`.

5.  **To Stop the Application**
    Press `Ctrl + C` in the terminal, then run the following command to shut down and remove the containers.
    ```bash
    docker-compose down
    ```

---

### 2. Running Manually (Local)

This method requires you to have Go and PostgreSQL installed and running on your system.

**Prerequisites:**

* [Go](https://golang.org/doc/install/) version 1.18 or newer.
* A running [PostgreSQL](https://www.postgresql.org/download/) instance.

**Steps:**

1.  **Clone the Repository**
    ```bash
    git clone https://github.com/lskeey/websocket-chat.git
    cd websocket-chat
    ```

2.  **Setup the Database**
    Connect to your PostgreSQL server and create a new database.
    ```sql
    CREATE DATABASE chat_app_db;
    ```

3.  **Configure Environment**
    Copy `.env.example` to `.env`.
    ```bash
    cp .env.example .env
    ```
    Open the `.env` file and adjust the `DB_` variables to match your local PostgreSQL configuration (host, port, user, password, and database name).

4.  **Install Go Dependencies**
    ```bash
    go mod tidy
    ```

5.  **Run the Application**
    This command will automatically run the database migrations and start the server.
    ```bash
    go run main.go
    ```
    The backend API is now running at `http://localhost:8080`.

## API Endpoints

Here is a list of the main available API endpoints:

* `POST /api/register` - Registers a new user.
* `POST /api/login` - Logs in and retrieves a JWT.
* `GET /api/users/search?username={query}` - (Protected) Searches for users.
* `GET /api/messages/{recipient_id}` - (Protected) Gets message history.
* `GET /api/ws` - (Protected) Upgrades the connection to a WebSocket for real-time chat.
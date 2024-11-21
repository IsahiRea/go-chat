# Real-Time Chat Application

A real-time chat system built with the following stack:
- **Frontend**: Javascript **(Work In Progress)**
- **Backend**: Go (with WebSockets)
- **Data Persistence**: MongoDB
- **Message Broadcasting**: Redis (Pub/Sub)
- **Authentication**: JWT (JSON Web Tokens)

Users can join chat rooms and exchange messages in real time. This project demonstrates the use of WebSocket connections, message persistence, Pub/Sub model, and real-time communication.

## Features
- Join multiple chat rooms
- Real-time messaging with WebSockets
- Redis Pub/Sub for broadcasting messages
- User authentication with JWT
- Message history stored in MongoDB
- Frontend

## Prerequisites
Before running the application, make sure you have the following installed:
- [Go](https://golang.org/)
- [Node.js](https://nodejs.org/) and [npm](https://www.npmjs.com/)
- [MongoDB](https://www.mongodb.com/)
- [Redis](https://redis.io/)

## Setup

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/go-chat.git
cd real-time-chat-app
```

### 2. Backend Setup (Go)
- **Navigate to the `backend/` directory**:
    ```bash
    cd backend
    ```
- **Install Go dependencies**:
    ```bash
    go mod tidy
    ```

- **Start Redis and MongoDB**:
    - Run Redis and MongoDB instances on your machine. If using Docker, you can use:
      ```bash
      docker run --name some-redis -p 6379:6379 -d redis
      docker run --name some-mongo -p 27017:27017 -d mongo
      ```

- **Run the Go server**:
    ```bash
    go run main.go
    ```

### 3. Frontend Setup
- **Navigate to the `frontend/` directory**:
    ```bash
    cd frontend
    ```

- **Install frontend dependencies**:
    ```bash
    npm install
    ```

- **Start the development server**:
    ```bash
    npm start
    ```

The application will run on `http://localhost:3000` by default.

## WebSocket API

### WebSocket Route
The WebSocket endpoint is available at:
```
ws://localhost:8080/ws?room={roomName}
```
- Replace `{roomName}` with the name of the chat room the user wants to join.

### Example WebSocket Message
```json
{
  "text": "Hello, everyone!",
  "user": "John Doe"
}
```

## Directory Structure
```
├── backend              # Go backend
│   ├── main.go          # Entry point of the Go backend
│   ├── handlers.go      # WebSocket and HTTP handlers
│   ├── auth.go          # JWT authentication logic
│   └── ...
├── frontend             
│   ├── src
│   ├── public
│   ├── temp
│   │   ├── index.html
│   │   ├── script.js
│   │   ├── styles.css
│   │   └── ...
│   └── ...
└── README.md            # Project documentation
```

## To-Do / Future Enhancements
- Add typing indicators
- Add file sharing functionality
- Improve user authentication (e.g., OAuth support)
- Add private messaging functionality
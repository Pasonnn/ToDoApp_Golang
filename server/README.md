# Todo API - Golang Backend

A RESTful API for managing todo tasks built with Golang and Gin framework.

## Features

- ✅ Create new tasks
- ✅ Get all tasks
- ✅ Update task completion status
- ✅ Delete tasks
- ✅ In-memory data store
- ✅ CORS enabled
- ✅ Proper error handling
- ✅ Consistent API response format

## API Endpoints

### 1. Create a new task
- **POST** `/tasks`
- **Request Body:**
```json
{
  "title": "Task title",
  "description": "Task description"
}
```

### 2. Get all tasks
- **GET** `/tasks`

### 3. Update task completion
- **PUT** `/tasks/{id}`
- **Request Body:**
```json
{
  "completed": true
}
```

### 4. Delete a task
- **DELETE** `/tasks/{id}`

## Response Format

All API responses follow this consistent format:

```json
{
  "success": true,
  "statusCode": 200,
  "message": "Descriptive message about the result",
  "data": {}
}
```

## Prerequisites

- Go 1.19 or higher

## Installation & Setup

1. **Clone the repository**
```bash
git clone <repository-url>
cd server
```

2. **Install dependencies**
```bash
go mod tidy
```

3. **Run the server**
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Environment Variables

- `PORT` - Server port (default: 8080)

## Project Structure

```
server/
├── main.go              # Application entry point
├── models/
│   └── task.go         # Data models and request/response structures
├── services/
│   └── task_service.go # Business logic and data operations
├── handlers/
│   └── task_handler.go # HTTP request handlers
├── routes/
│   └── routes.go       # Route configuration
├── go.mod              # Go module file
├── go.sum              # Dependency checksums
└── README.md           # This file
```

## Testing the API

### Using curl

1. **Create a task:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "description": "Study Go programming language"}'
```

2. **Get all tasks:**
```bash
curl http://localhost:8080/tasks
```

3. **Update task completion:**
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```

4. **Delete a task:**
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

### Using Postman

Import the following collection:

```json
{
  "info": {
    "name": "Todo API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Create Task",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"title\": \"Learn Go\",\n  \"description\": \"Study Go programming language\"\n}"
        },
        "url": {
          "raw": "http://localhost:8080/tasks",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["tasks"]
        }
      }
    },
    {
      "name": "Get All Tasks",
      "request": {
        "method": "GET",
        "url": {
          "raw": "http://localhost:8080/tasks",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["tasks"]
        }
      }
    },
    {
      "name": "Update Task",
      "request": {
        "method": "PUT",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"completed\": true\n}"
        },
        "url": {
          "raw": "http://localhost:8080/tasks/1",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["tasks", "1"]
        }
      }
    },
    {
      "name": "Delete Task",
      "request": {
        "method": "DELETE",
        "url": {
          "raw": "http://localhost:8080/tasks/1",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["tasks", "1"]
        }
      }
    }
  ]
}
```

## Deployment

### Docker Deployment

1. **Build the Docker image:**
```bash
docker build -t todo-api .
```

2. **Run the container:**
```bash
docker run -p 8080:8080 todo-api
```

### Cloud Deployment

#### Render
1. Connect your GitHub repository to Render
2. Create a new Web Service
3. Set build command: `go build -o main .`
4. Set start command: `./main`
5. Set environment variable: `PORT=8080`

#### Railway
1. Connect your GitHub repository to Railway
2. Railway will automatically detect Go and deploy
3. Set environment variable: `PORT=8080`

#### Fly.io
1. Install flyctl
2. Run: `fly launch`
3. Deploy: `fly deploy`

## Error Handling

The API returns appropriate HTTP status codes:

- `200` - Success
- `201` - Created
- `400` - Bad Request (invalid input)
- `404` - Not Found
- `500` - Internal Server Error

## CORS

CORS is enabled for all origins to allow frontend integration.

## License

This project is part of an intern developer test.

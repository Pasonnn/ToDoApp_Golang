# Docker Commands for Todo API

## 🐳 Container Management

### Start the container
```bash
docker run -d -p 8080:8080 --name todo-api-container todo-api
```

### Stop the container
```bash
docker stop todo-api-container
```

### Start an existing container
```bash
docker start todo-api-container
```

### Restart the container
```bash
docker restart todo-api-container
```

### Remove the container
```bash
docker rm -f todo-api-container
```

### View container logs
```bash
docker logs todo-api-container
```

### View container logs in real-time
```bash
docker logs -f todo-api-container
```

### Execute commands inside the container
```bash
docker exec -it todo-api-container sh
```

## 🏗️ Image Management

### Build the image
```bash
docker build -t todo-api .
```

### List all images
```bash
docker images
```

### Remove the image
```bash
docker rmi todo-api
```

### Remove all unused images
```bash
docker image prune -a
```

## 📊 Container Status

### List running containers
```bash
docker ps
```

### List all containers (including stopped)
```bash
docker ps -a
```

### Container resource usage
```bash
docker stats todo-api-container
```

## 🔧 Development Commands

### Build and run in one command
```bash
docker build -t todo-api . && docker run -d -p 8080:8080 --name todo-api-container todo-api
```

### Rebuild and restart
```bash
docker stop todo-api-container && docker rm todo-api-container && docker build -t todo-api . && docker run -d -p 8080:8080 --name todo-api-container todo-api
```

### Test the API
```bash
curl http://localhost:8080/tasks
```

## 🚀 Production Deployment

### Run with environment variables
```bash
docker run -d -p 8080:8080 -e PORT=8080 --name todo-api-container todo-api
```

### Run with custom port mapping
```bash
docker run -d -p 3000:8080 --name todo-api-container todo-api
```

### Run with restart policy
```bash
docker run -d -p 8080:8080 --restart unless-stopped --name todo-api-container todo-api
```

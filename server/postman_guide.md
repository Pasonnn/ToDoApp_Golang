# Postman Collection Guide for Todo API

## 📥 How to Import

### Method 1: Import Collection File
1. Open Postman
2. Click **Import** button (top left)
3. Select **Upload Files**
4. Choose `postman_collection.json`
5. Click **Import**

### Method 2: Import Environment File
1. Open Postman
2. Click **Import** button
3. Select **Upload Files**
4. Choose `postman_environment.json`
5. Click **Import**

## 🎯 Collection Overview

The collection includes **8 requests** covering all API endpoints:

### ✅ Success Cases
1. **Create Task** - POST `/tasks`
2. **Get All Tasks** - GET `/tasks`
3. **Update Task** - PUT `/tasks/{id}`
4. **Delete Task** - DELETE `/tasks/{id}`

### ❌ Error Cases
5. **Create Task - Invalid Request** - Missing required fields
6. **Update Task - Not Found** - Non-existent task
7. **Update Task - Invalid ID** - Invalid task ID format
8. **Delete Task - Not Found** - Non-existent task

## 🔧 Environment Setup

### 1. Select Environment
- Click the environment dropdown (top right)
- Select **"Todo API Environment"**

### 2. Environment Variables
- `base_url`: `http://localhost:8080` (default)
- `api_version`: `v1`
- `content_type`: `application/json`

### 3. Update for Production
When deploying to production, update the `base_url` variable:
- Render: `https://your-app-name.onrender.com`
- Railway: `https://your-app-name.railway.app`
- Fly.io: `https://your-app-name.fly.dev`

## 🧪 Testing Workflow

### Step 1: Start Your Server
```bash
# Local development
go run main.go

# Or using Docker
docker run -d -p 8080:8080 --name todo-api-container todo-api
```

### Step 2: Test Basic CRUD Operations
1. **Create Task** - Add a new task
2. **Get All Tasks** - Verify task was created
3. **Update Task** - Mark task as completed
4. **Get All Tasks** - Verify task was updated
5. **Delete Task** - Remove the task
6. **Get All Tasks** - Verify task was deleted

### Step 3: Test Error Cases
1. **Create Task - Invalid Request** - Test validation
2. **Update Task - Not Found** - Test 404 handling
3. **Update Task - Invalid ID** - Test bad request
4. **Delete Task - Not Found** - Test 404 handling

## ✅ Automated Tests

The collection includes **automated test scripts** that validate:

### Response Structure
- ✅ Status code is 200 or 201
- ✅ Response has required fields (`success`, `statusCode`, `message`, `data`)
- ✅ `success` field is boolean
- ✅ Status code matches response

### Manual Validation
Check these in the response:
- **Success responses** should have `"success": true`
- **Error responses** should have `"success": false`
- **Data field** should contain the actual payload
- **Message field** should be descriptive

## 📋 Example Test Sequence

### 1. Create Multiple Tasks
```json
// Task 1
{
  "title": "Learn Go",
  "description": "Study Go programming language"
}

// Task 2
{
  "title": "Build API",
  "description": "Create REST API with Gin framework"
}

// Task 3
{
  "title": "Deploy to Cloud",
  "description": "Deploy the application to cloud platform"
}
```

### 2. Verify All Tasks
- GET `/tasks` should return all 3 tasks

### 3. Update Tasks
- PUT `/tasks/1` with `{"completed": true}`
- PUT `/tasks/2` with `{"completed": true}`

### 4. Verify Updates
- GET `/tasks` should show tasks 1 and 2 as completed

### 5. Delete Tasks
- DELETE `/tasks/1`
- DELETE `/tasks/3`

### 6. Final Verification
- GET `/tasks` should return only task 2

## 🚀 Production Testing

### Update Environment
1. Change `base_url` to your production URL
2. Test all endpoints again
3. Verify CORS is working for frontend integration

### Common Production URLs
- **Render**: `https://your-app-name.onrender.com`
- **Railway**: `https://your-app-name.railway.app`
- **Fly.io**: `https://your-app-name.fly.dev`
- **Heroku**: `https://your-app-name.herokuapp.com`

## 🔍 Troubleshooting

### Connection Issues
- Verify server is running
- Check if port 8080 is available
- Ensure firewall allows connections

### CORS Issues
- Verify CORS is enabled in the backend
- Check if frontend domain is allowed

### Response Format Issues
- Ensure all responses follow the required format
- Check if `success`, `statusCode`, `message`, `data` fields are present

## 📊 Performance Testing

### Load Testing
- Use Postman's **Runner** feature
- Set iterations to test multiple requests
- Monitor response times

### Stress Testing
- Create many tasks rapidly
- Test concurrent requests
- Verify in-memory store handles load

## 🎉 Success Criteria

Your API is working correctly if:
- ✅ All success cases return proper JSON responses
- ✅ All error cases return appropriate status codes
- ✅ Response format is consistent across all endpoints
- ✅ CORS allows frontend integration
- ✅ Automated tests pass
- ✅ Manual validation confirms expected behavior

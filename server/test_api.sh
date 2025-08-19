#!/bin/bash

# Todo API Test Script
# This script tests all 4 API endpoints

BASE_URL="http://localhost:8080"

echo "🧪 Testing Todo API Endpoints"
echo "=============================="
echo ""

# Function to print colored output
print_result() {
    if [ $1 -eq 0 ]; then
        echo "✅ $2"
    else
        echo "❌ $2"
    fi
}

# Test 1: Create Task
echo "1. Testing POST /tasks - Create Task"
echo "-----------------------------------"
CREATE_RESPONSE=$(curl -s -X POST $BASE_URL/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Test Task", "description": "This is a test task"}')

echo "Response: $CREATE_RESPONSE"
TASK_ID=$(echo $CREATE_RESPONSE | grep -o '"id":[0-9]*' | cut -d':' -f2)
print_result $? "Create task"
echo ""

# Test 2: Get All Tasks
echo "2. Testing GET /tasks - Get All Tasks"
echo "------------------------------------"
GET_RESPONSE=$(curl -s $BASE_URL/tasks)
echo "Response: $GET_RESPONSE"
print_result $? "Get all tasks"
echo ""

# Test 3: Update Task
echo "3. Testing PUT /tasks/$TASK_ID - Update Task"
echo "--------------------------------------------"
UPDATE_RESPONSE=$(curl -s -X PUT $BASE_URL/tasks/$TASK_ID \
  -H "Content-Type: application/json" \
  -d '{"completed": true}')
echo "Response: $UPDATE_RESPONSE"
print_result $? "Update task"
echo ""

# Test 4: Delete Task
echo "4. Testing DELETE /tasks/$TASK_ID - Delete Task"
echo "-----------------------------------------------"
DELETE_RESPONSE=$(curl -s -X DELETE $BASE_URL/tasks/$TASK_ID)
echo "Response: $DELETE_RESPONSE"
print_result $? "Delete task"
echo ""

# Test 5: Verify Task Deleted
echo "5. Testing GET /tasks - Verify Task Deleted"
echo "-------------------------------------------"
FINAL_RESPONSE=$(curl -s $BASE_URL/tasks)
echo "Response: $FINAL_RESPONSE"
print_result $? "Verify task deleted"
echo ""

# Error Tests
echo "6. Testing Error Cases"
echo "----------------------"

# Test invalid request (missing description)
echo "Testing invalid request (missing description):"
ERROR_RESPONSE=$(curl -s -X POST $BASE_URL/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Incomplete Task"}')
echo "Response: $ERROR_RESPONSE"
print_result $? "Invalid request test"
echo ""

# Test non-existent task
echo "Testing non-existent task:"
NOT_FOUND_RESPONSE=$(curl -s -X GET $BASE_URL/tasks/999)
echo "Response: $NOT_FOUND_RESPONSE"
print_result $? "Not found test"
echo ""

echo "🎉 API Testing Complete!"
echo "========================"

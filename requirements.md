# **Intern Developer Test – Golang (Backend) + ReactJS (Frontend)**

## 🎯 Objective

Evaluate the candidate’s skills in **backend development (Golang)**, **frontend development (ReactJS/HTML+CSS+JS)**, and **cloud deployment**.

---

## 🔥 Project Requirements

### ✅ Requirement 1: Backend Development (To-do List API in Golang)

Build a **REST API** with the following features:

#### **Response Format (Required for ALL APIs)**

Every API must return a JSON response with this structure:

```json
{
  "success": true,
  "statusCode": 200,
  "message": "Descriptive message about the result",
  "data": {}
}
```

* `success`: `true` or `false`
* `statusCode`: HTTP status code (e.g., 200, 201, 400, 404, 500)
* `message`: short description of the operation result
* `data`: actual payload (object or array). For list endpoints, `data` should be an array.

---

#### **Endpoints**

1. **Create a new task**

   * **POST /tasks**
   * **Request Body**:

     ```json
     {
       "title": "Task title",
       "description": "Task description"
     }
     ```
   * **Response**:

     ```json
     {
       "success": true,
       "statusCode": 201,
       "message": "Task created successfully",
       "data": {
         "id": 1,
         "title": "Task title",
         "description": "Task description",
         "completed": false
       }
     }
     ```

2. **Get all tasks**

   * **GET /tasks**
   * **Response**:

     ```json
     {
       "success": true,
       "statusCode": 200,
       "message": "Tasks retrieved successfully",
       "data": [
         {
           "id": 1,
           "title": "Task title",
           "description": "Task description",
           "completed": false
         }
       ]
     }
     ```

3. **Update task completion**

   * **PUT /tasks/{id}**
   * **Request Body**:

     ```json
     {
       "completed": true
     }
     ```
   * **Response**:

     ```json
     {
       "success": true,
       "statusCode": 200,
       "message": "Task updated successfully",
       "data": {
         "id": 1,
         "title": "Task title",
         "description": "Task description",
         "completed": true
       }
     }
     ```

4. **Delete a task**

   * **DELETE /tasks/{id}**
   * **Response**:

     ```json
     {
       "success": true,
       "statusCode": 200,
       "message": "Task deleted successfully",
       "data": {
         "id": 1
       }
     }
     ```

---

#### **Technical Requirements**

* Golang with **Gin** or **Echo** framework.
* In-memory data store (`map` or `slice`).
* Proper error handling (e.g., invalid input → 400, not found → 404, internal error → 500).
* Organize code in layers (`routes/handlers`, `services`, `models`).

---

### ✅ Requirement 2: Frontend Development (ReactJS or Plain HTML+CSS+JS)

#### **Features**

1. Fetch and display tasks (`GET /tasks`).
2. Add a new task (`POST /tasks`).
3. Mark a task as completed (`PUT /tasks/{id}`).
4. Delete a task (`DELETE /tasks/{id}`).
5. Responsive UI (mobile-friendly, simple and clean).

#### **Technical**

* ReactJS (preferred) or vanilla HTML+CSS+JS.
* API calls via `axios` or `fetch`.
* Reuse components (`TaskList`, `TaskItem`, `AddTaskForm`).
* Show loading/error states for API calls.

---

### ✅ Requirement 3: Deployment

1. Choose a **free-tier cloud**:

   * Backend → Render / Railway / Fly.io / AWS Free Tier
   * Frontend → Vercel / Netlify

2. Deploy **backend API**.

3. Deploy **frontend** (configured to call deployed backend).

4. Enable **CORS** for API.

5. Add `README.md` with setup + deployment instructions.

6. (Optional but preferred) Package backend in **Docker**.

---

## 📋 Submission

1. Requirement 1 → GitHub repo (backend).
2. Requirement 1 + 2 → GitHub repo (backend + frontend).
3. Requirement 1 + 3 → GitHub repo + Postman collection for testing APIs.
4. Full (1 + 2 + 3) → GitHub repo + working deployment links + Postman collection.

---

## 📅 Deadline

**7 days** after receiving the test.

---

## 🔑 Evaluation Criteria

* Clean, readable code.
* Consistent response format (`success`, `statusCode`, `message`, `data`).
* Proper error handling.
* Successful deployment (frontend + backend works together).
* UI/UX quality.

---
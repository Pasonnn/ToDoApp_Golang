package services

import (
	"errors"
	"sync"
	"time"

	"todo-api/models"
)

// TaskService handles business logic for tasks
type TaskService struct {
	tasks map[int]*models.Task
	mutex sync.RWMutex
	nextID int
}

// NewTaskService creates a new task service instance
func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[int]*models.Task),
		nextID: 1,
	}
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(req *models.CreateTaskRequest) (*models.Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task := &models.Task{
		ID:          s.nextID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tasks[task.ID] = task
	s.nextID++

	return task, nil
}

// GetAllTasks returns all tasks
func (s *TaskService) GetAllTasks() []*models.Task {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	tasks := make([]*models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

// GetTaskByID returns a task by ID
func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return task, nil
}

// UpdateTask updates a task's completion status
func (s *TaskService) UpdateTask(id int, req *models.UpdateTaskRequest) (*models.Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	task.Completed = req.Completed
	task.UpdatedAt = time.Now()

	return task, nil
}

// DeleteTask deletes a task by ID
func (s *TaskService) DeleteTask(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}

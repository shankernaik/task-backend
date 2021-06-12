package services

import (
	"errors"

	"github.com/shankernaik/task-backend/task-backend/models"
)

// TaskService is a service struct
type TaskService struct {
}

// NewService creates a new TaskService
func NewService() *TaskService {
	return &TaskService{}
}

var tasks []models.Task

func init() {
	tasks = append(tasks, models.Task{ID: "1", Title: "My first task", Description: "This is the content of my first task"})
}

// GetTasks service returns all the tasks
func (svc *TaskService) GetTasks() []models.Task {
	return tasks
}

// CreateTask service adds a new task
func (svc *TaskService) CreateTask(task models.Task) []models.Task {
	tasks = append(tasks, task)
	return tasks
}

// GetTask service returns a task
func (svc *TaskService) GetTask(id string) (models.Task, error) {

	for _, item := range tasks {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Task{}, errors.New("Not found")
}

// UpdateTask service updates a task
func (svc *TaskService) UpdateTask(task models.Task) (models.Task, error) {

	for index := range tasks {
		if tasks[index].ID == task.ID {
			tasks[index].Title = task.Title
			tasks[index].Description = task.Description
			return tasks[index], nil
		}
	}
	return models.Task{}, errors.New("Not found")
}

// DeleteTask service deletes a task
func (svc *TaskService) DeleteTask(id string) ([]models.Task, error) {

	for index, item := range tasks {
		if item.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	return tasks, nil
}

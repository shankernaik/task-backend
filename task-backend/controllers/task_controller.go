package controller

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/shankernaik/task-backend/task-backend/models"
	"github.com/shankernaik/task-backend/task-backend/services"
)

// TaskController is a controller struct
type TaskController struct {
	TaskSvc *services.TaskService
}

// NewController returns a TaskController pointer
func NewController(taskTrackerSvc *services.TaskService) *TaskController {
	return &TaskController{TaskSvc: taskTrackerSvc}
}

// GetTasks returns all the tasks
func (ttc *TaskController) GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, ttc.TaskSvc.GetTasks())
}

// CreateTask adds a new task
func (ttc *TaskController) CreateTask(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task.ID = strconv.Itoa(rand.Intn(1000000))
	return c.JSON(http.StatusOK, ttc.TaskSvc.CreateTask(task))
}

// GetTask returns a task
func (ttc *TaskController) GetTask(c echo.Context) error {
	id := c.Param("id")

	task, err := ttc.TaskSvc.GetTask(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

// UpdateTask updates a task
func (ttc *TaskController) UpdateTask(c echo.Context) error {

	var task models.Task

	if err := c.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task, err := ttc.TaskSvc.UpdateTask(task)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task
func (ttc *TaskController) DeleteTask(c echo.Context) error {

	id := c.Param("id")

	tasks, err := ttc.TaskSvc.DeleteTask(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

package routes

import (
	"github.com/labstack/echo"
	controller "github.com/shankernaik/task-backend/task-backend/controllers"
	"github.com/shankernaik/task-backend/task-backend/services"
)

//Routes configures all the http routes and their handlers
func Routes(e *echo.Echo) {
	TaskController := controller.NewController(services.NewService())
	e.GET("/api/v1/tasks", TaskController.GetTasks)
	e.POST("/api/v1/tasks", TaskController.CreateTask)
	e.GET("/api/v1/tasks/:id", TaskController.GetTask)
	e.PUT("/api/v1/tasks/", TaskController.UpdateTask)
	e.DELETE("/api/v1/tasks/:id", TaskController.DeleteTask)
}

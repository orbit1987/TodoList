package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/orbit1987/TodoList/internal/api/http/v1/domain"
	"net/http"
)

func (handler *Handler) createTask(c echo.Context) error {
	newTask := new(domain.NewTask)
	if err := c.Bind(newTask); err != nil {
		response := domain.ResponseMess{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	id, err := handler.services.TodoTask.CreateTask(
		newTask.Name,
		newTask.Description,
		newTask.Status,
	)

	if err != nil {
		response := domain.ResponseMess{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := domain.ResponseIdAndMess{
		Id:      id,
		Message: "success",
	}

	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) updateTask(c echo.Context) error {
	updateTask := new(domain.UpdateTask)
	if err := c.Bind(updateTask); err != nil {
		response := domain.ResponseMess{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	id, err := handler.services.TodoTask.UpdateTask(
		updateTask.Id,
		updateTask.Name,
		updateTask.Description,
		updateTask.Status,
	)

	if err != nil {
		response := domain.ResponseMess{Message: err.Error()}
		return c.JSON(http.StatusOK, response)
	}

	response := domain.ResponseIdAndMess{Message: "success", Id: id}
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) deleteTask(c echo.Context) error {
	id := c.Param("id")
	err := handler.services.TodoTask.DeleteTask(id)
	if err != nil {
		taskResponse := domain.ResponseMess{Message: err.Error()}
		return c.JSON(http.StatusOK, taskResponse)
	}

	taskResponse := domain.ResponseMess{}
	taskResponse.Message = "success"
	return c.JSON(http.StatusOK, taskResponse)
}

func (handler *Handler) getTaskItem(c echo.Context) error {
	id := c.Param("id")
	task, err := handler.services.TodoTask.GetTaskById(id)
	if err != nil {
		taskResponse := domain.ResponseMess{Message: err.Error()}
		return c.JSON(http.StatusOK, taskResponse)
	}

	fullTaskById := new(domain.FullTask)
	fullTaskById.Id = task.Id
	fullTaskById.Description = task.Description
	fullTaskById.Status = task.Status
	fullTaskById.TimeStump = task.TimeStump

	return c.JSON(http.StatusOK, task)
}

func (handler *Handler) taskList(c echo.Context) error {
	tasks := handler.services.TodoTask.GetTaskList()
	tasksResponse := domain.AllTasks{}
	for _, task := range tasks {
		tasksResponse.AllTasks = append(tasksResponse.AllTasks, domain.FullTask{
			Id:          task.Id,
			Name:        task.Name,
			Description: task.Description,
			Status:      task.Status,
			TimeStump:   task.TimeStump,
		})
	}

	return c.JSON(http.StatusOK, tasksResponse)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "api/v1.0")
		return next(c)
	}
}

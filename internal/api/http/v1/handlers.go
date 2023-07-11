package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) createTask(c echo.Context) error {
	newTask := NewTask{}

	if err := c.Bind(&newTask); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	id, err := handler.services.TodoTask.CreateTask(
		newTask.Name,
		newTask.Description,
		newTask.Status,
	)
	if err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	response := ResponseIdAndMess{Id: id, Message: "success"}
	return response.sendResponse(c, http.StatusOK)
}

func (handler *Handler) updateTask(c echo.Context) error {
	id := c.Param("id")

	updateData := make(map[string]interface{})
	if err := c.Bind(&updateData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	id, err := handler.services.TodoTask.UpdateTask(id, updateData)
	if err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusOK)
	}

	response := ResponseIdAndMess{Message: "success", Id: id}
	return response.sendResponse(c, http.StatusOK)
}

func (handler *Handler) deleteTask(c echo.Context) error {
	id := c.Param("id")
	taskResponse := ResponseMess{}

	err := handler.services.TodoTask.DeleteTask(id)
	if err != nil {
		taskResponse.Message = err.Error()
	} else {
		taskResponse.Message = "success"
	}

	return taskResponse.sendResponse(c, http.StatusOK)
}

func (handler *Handler) getTaskItem(c echo.Context) error {
	id := c.Param("id")

	task, err := handler.services.TodoTask.GetTaskById(id)
	if err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusOK)
	}

	fullTaskById := new(ResponseFullTask)
	fullTaskById.Id = task.Id
	fullTaskById.Description = task.Description
	fullTaskById.Status = task.Status
	fullTaskById.TimeStump = task.TimeStump

	return c.JSON(http.StatusOK, task)
}

func (handler *Handler) taskList(c echo.Context) error {
	tasks := handler.services.TodoTask.GetTaskList()

	tasksResponse := ResponseAllTasks{}
	for _, task := range tasks {
		tasksResponse.AllTasks = append(tasksResponse.AllTasks, ResponseFullTask{
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

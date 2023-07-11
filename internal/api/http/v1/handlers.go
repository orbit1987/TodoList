package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/orbit1987/TodoList/internal/tools"
	"net/http"
)

func (handler *Handler) createTask(c echo.Context) error {
	newTaskData := make(map[string]interface{})
	if err := c.Bind(&newTaskData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	newTask := NewTask{}
	if name, ok := newTaskData[name].(string); ok {
		newTask.Name = name
	}

	if description, ok := newTaskData[description].(string); ok {
		newTask.Description = description
	}

	if status, ok := newTaskData[status].(float64); ok {
		newTask.Status = int(status)
	}

	if tools.MapHasOnlyValidData(newTaskData, name, description, status) {
		response := ResponseMess{Message: "invalid json body"}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if len(newTaskData) == 0 {
		message := fmt.Sprintf("can not update %s json body is empty", taskId)
		response := ResponseMess{Message: message}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if len(newTask.Name) == 0 {
		message := errors.New("new task could not be created, field name must be required").Error()
		response := ResponseMess{Message: message}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if newTask.Status != 0 && newTask.Status != 1 {
		message := fmt.Sprintf("new task could not be created, status value can be %d or %d", 0, 1)
		response := ResponseMess{Message: errors.New(message).Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	taskId, err := handler.services.TodoTask.CreateTask(
		newTask.Name,
		newTask.Description,
		newTask.Status,
	)
	if err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusInternalServerError)
	}

	response := ResponseIdAndMess{TaskId: taskId, Message: success}
	return response.sendResponse(c, http.StatusOK)
}

func (handler *Handler) updateTask(c echo.Context) error {
	id := c.Param(taskId)

	updateData := make(map[string]interface{})
	if err := c.Bind(&updateData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	delete(updateData, taskId)
	if len(updateData) == 0 {
		message := fmt.Sprintf("can not update %s json body is empty", taskId)
		response := ResponseMess{Message: message}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if tools.MapHasOnlyValidData(updateData, name, description, status) {
		response := ResponseMess{Message: "invalid json body"}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	message := fmt.Sprintf("new task could not be created, status value can be %d or %d", 0, 1)
	if status, ok := updateData[status].(float64); ok {
		taskStatus := int(status)
		if taskStatus != 0 && taskStatus != 1 {
			response := ResponseMess{Message: errors.New(message).Error()}
			return response.sendResponse(c, http.StatusBadRequest)
		}
	}

	taskId, err := handler.services.TodoTask.UpdateTask(id, updateData)
	if err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusNotFound)
	}

	response := ResponseIdAndMess{Message: success, TaskId: taskId}
	return response.sendResponse(c, http.StatusOK)
}

func (handler *Handler) deleteTask(c echo.Context) error {
	taskResponse := ResponseMess{}
	err := handler.services.TodoTask.DeleteTask(c.Param(taskId))
	if err != nil {
		taskResponse.Message = err.Error()
		return taskResponse.sendResponse(c, http.StatusNotFound)
	}

	taskResponse.Message = success
	return taskResponse.sendResponse(c, http.StatusOK)
}

func (handler *Handler) getTaskItem(c echo.Context) error {
	task, err := handler.services.TodoTask.GetTaskById(c.Param(taskId))
	if err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusNotFound)
	}

	fullTaskById := new(ResponseFullTask)
	fullTaskById.TaskId = task.TaskId
	fullTaskById.Name = task.Name
	fullTaskById.Description = task.Description
	fullTaskById.Status = task.Status
	fullTaskById.TimeStump = task.TimeStump

	return c.JSON(http.StatusOK, fullTaskById)
}

func (handler *Handler) taskList(c echo.Context) error {
	tasks := handler.services.TodoTask.GetTaskList()
	tasksResponse := ResponseAllTasks{}
	for _, task := range tasks {
		tasksResponse.AllTasks = append(tasksResponse.AllTasks, ResponseFullTask{
			TaskId:      task.TaskId,
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

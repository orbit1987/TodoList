package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *Handler) createTask(c echo.Context) error {
	userToken := c.Request().Header.Get(token)
	newTaskData := make(map[string]interface{})

	if err := c.Bind(&newTaskData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if err := MapHasOnlyValidData(newTaskData, name, description, status); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if err := mapIsNotEmpty(newTaskData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	newTask := NewTask{}
	newTask.fill(newTaskData)

	if err := newTask.nameValidation(); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if err := newTask.statusValidation(); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	taskId, err := handler.services.TodoTask.CreateTask(
		userToken,
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
	userToken := c.Request().Header.Get(token)
	id := c.Param(taskId)

	updateData := make(map[string]interface{})
	if err := c.Bind(&updateData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	delete(updateData, taskId)
	if err := mapIsNotEmpty(updateData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if err := MapHasOnlyValidData(updateData, name, description, status); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	if err := statusValidation(updateData); err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusBadRequest)
	}

	taskId, err := handler.services.TodoTask.UpdateTask(userToken, id, updateData)
	if err != nil {
		response := ResponseMess{Message: err.Error()}
		return response.sendResponse(c, http.StatusNotFound)
	}

	response := ResponseIdAndMess{Message: success, TaskId: taskId}
	return response.sendResponse(c, http.StatusOK)
}

func (handler *Handler) deleteTask(c echo.Context) error {
	userToken := c.Request().Header.Get(token)

	taskResponse := ResponseMess{}
	err := handler.services.TodoTask.DeleteTask(userToken, c.Param(taskId))
	if err != nil {
		taskResponse.Message = err.Error()
		return taskResponse.sendResponse(c, http.StatusNotFound)
	}

	taskResponse.Message = success
	return taskResponse.sendResponse(c, http.StatusOK)
}

func (handler *Handler) deleteUserTaskList(c echo.Context) error {
	userToken := c.Request().Header.Get(token)

	taskResponse := ResponseMess{}
	err := handler.services.TodoTask.DeleteUserTaskList(userToken)
	if err != nil {
		taskResponse.Message = err.Error()
		return taskResponse.sendResponse(c, http.StatusNotFound)
	}

	taskResponse.Message = success
	return taskResponse.sendResponse(c, http.StatusOK)
}

func (handler *Handler) getTaskItem(c echo.Context) error {
	userToken := c.Request().Header.Get(token)
	task, err := handler.services.TodoTask.GetTaskById(userToken, c.Param(taskId))
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

func (handler *Handler) getUserTasksList(c echo.Context) error {
	userToken := c.Request().Header.Get(token)

	tasks := handler.services.TodoTask.GetUserTaskList(userToken)
	tasksResponse := ResponseAllUserTasks{}
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

func (handler *Handler) taskList(c echo.Context) error {
	usersMap := handler.services.TodoTask.GetTaskList()
	base := ResponseAllData{Users: make(map[string]ResponseAllUserTasks)}
	for user, userBase := range usersMap {
		tasksResponse := ResponseAllUserTasks{}
		for _, task := range userBase {
			tasksResponse.AllTasks = append(tasksResponse.AllTasks, ResponseFullTask{
				TaskId:      task.TaskId,
				Name:        task.Name,
				Description: task.Description,
				Status:      task.Status,
				TimeStump:   task.TimeStump,
			})
		}
		base.Users[user] = tasksResponse
	}

	return c.JSON(http.StatusOK, base)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "api/v1.0")
		return next(c)
	}
}

func checkUserToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Request().Header.Get(token)
		if len(userToken) == 0 {
			response := ResponseMess{Message: "user token required"}
			return response.sendResponse(c, http.StatusBadRequest)
		}

		return next(c)
	}
}

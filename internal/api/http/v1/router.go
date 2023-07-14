package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/orbit1987/TodoList/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (handler *Handler) InitRouter() http.Handler {
	router := echo.New()
	api := router.Group(fmt.Sprintf("/%s/%s", api, v1), serverHeader)

	api.POST("/createTask", handler.createTask, checkUserToken)
	api.PUT(fmt.Sprintf("/updateTask/:%s", taskId), handler.updateTask, checkUserToken)

	api.DELETE(fmt.Sprintf("/deleteTask/:%s", taskId), handler.deleteTask, checkUserToken)
	api.DELETE(fmt.Sprintf("/deleteUserTaskList"), handler.deleteUserTaskList, checkUserToken)

	api.GET(fmt.Sprintf("/getTaskItem/:%s", taskId), handler.getTaskItem, checkUserToken)
	api.GET("/getUserTasksList", handler.getUserTaskList, checkUserToken)
	api.GET("/tasksList", handler.taskList)

	return router
}

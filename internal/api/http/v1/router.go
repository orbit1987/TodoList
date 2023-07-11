package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/orbit1987/TodoList/internal/service"
	"net/http"
)

const (
	api = "api"
	v1  = "v1"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (handler *Handler) InitRouter() http.Handler {
	router := echo.New()
	api := router.Group(getRootPath(), serverHeader)
	api.POST("/createTask", handler.createTask)
	api.PUT("/updateTask", handler.updateTask)
	api.DELETE("/deleteTask/:id", handler.deleteTask)
	api.GET("/tasksList", handler.taskList)
	api.GET("/getTaskItem/:id", handler.getTaskItem)

	return router
}

func getRootPath() string {
	return fmt.Sprintf("/%s/%s", api, v1)
}

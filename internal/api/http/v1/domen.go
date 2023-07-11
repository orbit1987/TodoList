package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type NewTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type UpdateTask struct {
	Id          string `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type TaskId struct {
	Id string `json:"Id"`
}

type ResponseFullTask struct {
	Id          string `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	TimeStump   int64  `json:"timeStump"`
}

type ResponseAllTasks struct {
	AllTasks []ResponseFullTask `json:"tasks"`
}

type ResponseMess struct {
	Message string `json:"message"`
}

type ResponseIdAndMess struct {
	Id      string `json:"taskId"`
	Message string `json:"message"`
}

func (response *ResponseMess) sendResponse(c echo.Context, statusCode int) error {
	fmt.Println(response.Message)
	return c.JSON(statusCode, response)
}

func (response *ResponseIdAndMess) sendResponse(c echo.Context, statusCode int) error {
	fmt.Println(response.Id, response.Message)
	return c.JSON(statusCode, response)
}
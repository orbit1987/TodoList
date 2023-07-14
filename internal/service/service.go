package service

import (
	"github.com/orbit1987/TodoList/internal/domain"
	"github.com/orbit1987/TodoList/internal/repository"
)

type TodoTask interface {
	CreateTask(user string, name string, description string, status int) (string, error)
	UpdateTask(userToken string, taskId string, updateData map[string]interface{}) (string, error)

	DeleteTask(userToken string, taskId string) error
	DeleteUserTaskList(userToken string) error

	GetTaskById(userToken string, taskId string) (*domain.Task, error)
	GetUserTaskList(userToken string) map[string]*domain.Task
	GetTaskList() map[string]map[string]*domain.Task
}

type Service struct {
	TodoTask TodoTask
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		TodoTask: NewTaskService(repository),
	}
}

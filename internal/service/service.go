package service

import (
	"github.com/orbit1987/TodoList/internal/domain"
	"github.com/orbit1987/TodoList/internal/repository"
)

type TodoTask interface {
	CreateTask(name string, description string, status int) (string, error)
	UpdateTask(taskId string, updateData map[string]interface{}) (string, error)
	DeleteTask(taskId string) error
	GetTaskList() map[string]*domain.Task
	GetTaskById(taskId string) (task *domain.Task, err error)
}

type Service struct {
	TodoTask TodoTask
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		TodoTask: NewTaskService(repository),
	}
}

package repository

import "github.com/orbit1987/TodoList/internal/domain"

type TodoTask interface {
	CreateTask(task *domain.Task) (string, error)
	UpdateTask(taskId string, updateData map[string]interface{}) (string, error)
	DeleteTask(taskId string) error
	GetTaskList() map[string]*domain.Task
	GetTaskById(taskId string) (*domain.Task, error)
}

type Repository struct {
	TodoTask TodoTask
}

func NewRepository() *Repository {
	base := make(map[string]*domain.Task)
	return &Repository{
		TodoTask: NewTaskBase(base),
	}
}

package repository

import "github.com/orbit1987/TodoList/internal/domain"

type TodoTask interface {
	CreateTask(task *domain.Task) (id string, err error)
	UpdateTask(task *domain.Task)
	DeleteTask(taskId string) error
	GetTaskList() map[string]*domain.Task
	GetTaskById(taskId string) (task *domain.Task, err error)
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

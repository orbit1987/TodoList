package repository

import "github.com/orbit1987/TodoList/internal/domain"

type TodoTask interface {
	CreateTask(userToken string, task *domain.Task) (string, error)
	UpdateTask(userToken string, taskId string, updateData map[string]interface{}) (string, error)

	DeleteTask(userToken string, taskId string) error
	DeleteUserTaskList(userToken string) error

	GetTaskById(userToken string, taskId string) (*domain.Task, error)
	GetUserTaskList(userToken string) map[string]*domain.Task
	GetTaskList() map[string]map[string]*domain.Task
}

type Repository struct {
	TodoTask TodoTask
}

func NewRepository() *Repository {
	dataBase := make(map[string]map[string]*domain.Task)
	return &Repository{
		TodoTask: NewTaskBase(dataBase),
	}
}

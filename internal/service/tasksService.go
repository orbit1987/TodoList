package service

import (
	"github.com/google/uuid"
	"github.com/orbit1987/TodoList/internal/domain"
	"github.com/orbit1987/TodoList/internal/repository"
	"time"
)

type TaskService struct {
	repository *repository.Repository
}

func NewTaskService(repository *repository.Repository) *TaskService {
	return &TaskService{repository: repository}
}

func (t *TaskService) CreateTask(name string, description string, status int) (id string, err error) {
	newTask := new(domain.Task)
	newTask.Id = uuid.New().String()
	newTask.Name = name
	newTask.Description = description
	newTask.Status = status
	newTask.TimeStump = time.Now().UnixMilli()

	return t.repository.TodoTask.CreateTask(newTask)
}

func (t *TaskService) UpdateTask(taskId string, name string, description string, status int) (id string, err error) {
	task, err := t.repository.TodoTask.GetTaskById(taskId)
	if err != nil {
		return "", err
	}

	task.Name = name
	task.Description = description
	task.Status = status

	return task.Id, nil
}

func (t *TaskService) DeleteTask(taskId string) error {
	return t.repository.TodoTask.DeleteTask(taskId)
}

func (t *TaskService) GetTaskList() map[string]*domain.Task {
	return t.repository.TodoTask.GetTaskList()
}

func (t *TaskService) GetTaskById(taskId string) (task *domain.Task, err error) {
	return t.repository.TodoTask.GetTaskById(taskId)
}

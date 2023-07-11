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

func (t *TaskService) UpdateTask(taskId string, updateData map[string]interface{}) (id string, err error) {
	task, error := t.repository.TodoTask.GetTaskById(taskId)
	if error != nil {
		return "", error
	}

	if name, ok := updateData["name"].(string); ok {
		task.Name = name
	}

	if description, ok := updateData["description"].(string); ok {
		task.Description = description
	}

	if status, ok := updateData["status"].(float64); ok {
		task.Status = int(status)
	}

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

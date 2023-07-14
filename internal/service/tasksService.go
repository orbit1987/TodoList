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

func (t *TaskService) CreateTask(userToken string, name string, description string, status int) (string, error) {
	newTask := new(domain.Task)
	newTask.TaskId = uuid.New().String()
	newTask.Name = name
	newTask.Description = description
	newTask.Status = status
	newTask.TimeStump = time.Now().UnixMilli()

	return t.repository.TodoTask.CreateTask(userToken, newTask)
}

func (t *TaskService) UpdateTask(userToken string, taskId string, updateData map[string]interface{}) (string, error) {
	return t.repository.TodoTask.UpdateTask(userToken, taskId, updateData)
}

func (t *TaskService) DeleteTask(userToken string, taskId string) error {
	return t.repository.TodoTask.DeleteTask(userToken, taskId)
}

func (t *TaskService) DeleteUserTaskList(userToken string) error {
	return t.repository.TodoTask.DeleteUserTaskList(userToken)
}

func (t *TaskService) GetTaskById(userToken string, taskId string) (*domain.Task, error) {
	return t.repository.TodoTask.GetTaskById(userToken, taskId)
}

func (t *TaskService) GetUserTaskList(userToken string) map[string]*domain.Task {
	userTaskList := t.repository.TodoTask.GetUserTaskList(userToken)
	if userTaskList == nil {
		userTaskList = make(map[string]*domain.Task)
	}

	return userTaskList
}

func (t *TaskService) GetTaskList() map[string]map[string]*domain.Task {
	return t.repository.TodoTask.GetTaskList()
}

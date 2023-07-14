package repository

import (
	"errors"
	"fmt"
	"github.com/orbit1987/TodoList/internal/domain"
)

type TaskRepository struct {
	dataBase map[string]map[string]*domain.Task
}

func NewTaskBase(dataBase map[string]map[string]*domain.Task) *TaskRepository {
	return &TaskRepository{dataBase: dataBase}
}

func (base *TaskRepository) CreateTask(userToken string, task *domain.Task) (string, error) {
	dataData := base.dataBase[userToken]
	if dataData == nil {
		dataData = make(map[string]*domain.Task)
		base.dataBase[userToken] = dataData
	}
	dataData[task.TaskId] = task

	return task.TaskId, nil
}

func (base *TaskRepository) UpdateTask(userToken string, taskId string, updateData map[string]interface{}) (string, error) {
	task, err := base.getUserTask(userToken, taskId)
	if err != nil {
		return "", err
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

	return task.TaskId, nil
}

func (base *TaskRepository) DeleteTask(userToken string, taskId string) error {
	_, err := base.getUserTask(userToken, taskId)
	if err != nil {
		return err
	}

	delete(base.dataBase[userToken], taskId)
	if len(base.dataBase[userToken]) == 0 {
		delete(base.dataBase, userToken)
	}

	return nil
}

func (base *TaskRepository) DeleteUserTaskList(userToken string) error {
	dataBase := base.dataBase[userToken]
	if dataBase == nil {
		return tokenNotFoundError(userToken)
	}

	delete(base.dataBase, userToken)

	return nil
}

func (base *TaskRepository) GetTaskById(userToken string, taskId string) (*domain.Task, error) {
	return base.getUserTask(userToken, taskId)
}

func (base *TaskRepository) GetUserTaskList(userToken string) map[string]*domain.Task {
	return base.dataBase[userToken]
}

func (base *TaskRepository) GetTaskList() map[string]map[string]*domain.Task {
	return base.dataBase
}

func (base *TaskRepository) getUserTask(userToken string, taskId string) (*domain.Task, error) {
	dataBase := base.dataBase[userToken]
	if dataBase == nil {
		return &domain.Task{}, tokenNotFoundError(userToken)
	}

	task := dataBase[taskId]
	if task == nil {
		return &domain.Task{}, taskIdNotFoundError(taskId, userToken)
	}

	return task, nil
}

func taskIdNotFoundError(taskId string, userToken string) error {
	return errors.New(fmt.Sprintf("taskId %s for userToken %s not found", taskId, userToken))
}

func tokenNotFoundError(userToken string) error {
	return errors.New(fmt.Sprintf("userToken %s not found", userToken))
}

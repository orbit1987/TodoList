package repository

import (
	"errors"
	"fmt"
	"github.com/orbit1987/TodoList/internal/domain"
)

type TaskRepository struct {
	base map[string]*domain.Task
}

func NewTaskBase(base map[string]*domain.Task) *TaskRepository {
	return &TaskRepository{base: base}
}

func (taskBase *TaskRepository) CreateTask(task *domain.Task) (string, error) {
	taskBase.base[task.TaskId] = task
	return task.TaskId, nil
}

func (taskBase *TaskRepository) UpdateTask(taskId string, updateData map[string]interface{}) (string, error) {
	task := taskBase.base[taskId]
	if task == nil {
		return "", taskIdNotFoundError(taskId)
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

func (taskBase *TaskRepository) DeleteTask(taskId string) error {
	lengthBeforeDelete := len(taskBase.base)
	delete(taskBase.base, taskId)
	lengthAfterDelete := len(taskBase.base)

	if lengthBeforeDelete == lengthAfterDelete {
		return taskIdNotFoundError(taskId)
	}

	return nil
}

func (taskBase *TaskRepository) GetTaskList() map[string]*domain.Task {
	return taskBase.base
}

func (taskBase *TaskRepository) GetTaskById(taskId string) (*domain.Task, error) {
	task := taskBase.base[taskId]
	if task == nil {
		return nil, taskIdNotFoundError(taskId)
	}

	return task, nil
}

func taskIdNotFoundError(taskId string) error {
	return errors.New(fmt.Sprintf("taskId %s not found", taskId))
}

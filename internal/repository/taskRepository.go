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

func (taskBase *TaskRepository) CreateTask(task *domain.Task) (id string, err error) {
	taskBase.base[task.Id] = task
	return task.Id, nil
}

func (taskBase *TaskRepository) UpdateTask(task *domain.Task) {
	taskBase.base[task.Id] = task
}

func (taskBase *TaskRepository) DeleteTask(taskId string) error {
	lengthBeforeDelete := len(taskBase.base)
	delete(taskBase.base, taskId)
	lengthAfterDelete := len(taskBase.base)

	if lengthBeforeDelete == lengthAfterDelete {
		return errors.New(fmt.Sprintf("taskId %s not found", taskId))
	}

	return nil
}

func (taskBase *TaskRepository) GetTaskList() map[string]*domain.Task {
	return taskBase.base
}

func (taskBase *TaskRepository) GetTaskById(taskId string) (*domain.Task, error) {
	task := taskBase.base[taskId]
	if task == nil {
		return nil, errors.New(fmt.Sprintf("taskId %s not found", taskId))
	}

	return task, nil
}

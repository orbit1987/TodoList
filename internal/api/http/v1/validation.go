package handler

import (
	"errors"
	"fmt"
)

func (newTask *NewTask) fill(newTaskData map[string]interface{}) {
	if name, ok := newTaskData[name].(string); ok {
		newTask.Name = name
	}

	if description, ok := newTaskData[description].(string); ok {
		newTask.Description = description
	}

	if status, ok := newTaskData[status].(float64); ok {
		newTask.Status = int(status)
	}
}

func (newTask *NewTask) nameValidation() error {
	if len(newTask.Name) == 0 {
		return errors.New("new task could not be created, field name must be required")
	}

	return nil
}

func (newTask *NewTask) statusValidation() error {
	if newTask.Status != 0 && newTask.Status != 1 {
		message := fmt.Sprintf("new task could not be created, status value can be %d or %d", 0, 1)
		return errors.New(message)
	}

	return nil
}

func mapIsNotEmpty(updateData map[string]interface{}) error {
	if len(updateData) == 0 {
		message := fmt.Sprintf("json body is empty")
		return errors.New(message)
	}

	return nil
}

func MapHasOnlyValidData(data map[string]interface{}, validation ...string) error {
	validationMap := make(map[string]interface{})
	for k, v := range data {
		validationMap[k] = v
	}

	for _, v := range validation {
		delete(validationMap, v)
	}

	if len(validationMap) > 0 {
		return errors.New("invalid json body")
	}

	return nil
}

func statusValidation(updateData map[string]interface{}) error {
	message := fmt.Sprintf("new task could not be created, status value can be %d or %d", 0, 1)
	if status, ok := updateData[status].(float64); ok {
		taskStatus := int(status)
		if taskStatus != 0 && taskStatus != 1 {
			return errors.New(message)
		}
	}

	return nil
}

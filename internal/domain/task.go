package domain

type Status string

const (
	completed    Status = "completed"
	notCompleted Status = "notCompleted"
)

type Task struct {
	id          string
	name        string
	description string
	status      Status
}

package domain

type Status string

type NewTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type UpdateTask struct {
	Id          string `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type TaskId struct {
	Id string `json:"Id"`
}

type FullTask struct {
	Id          string `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	TimeStump   int64  `json:"timeStump"`
}

type AllTasks struct {
	AllTasks []FullTask `json:"allTasks"`
}

type ResponseIdAndMess struct {
	Id      string `json:"taskId"`
	Message string `json:"message"`
}

type ResponseMess struct {
	Message string `json:"message"`
}

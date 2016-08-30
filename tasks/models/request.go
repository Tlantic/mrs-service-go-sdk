package models

type TasksRequest struct {
	Status string `json:"status"`
	Result []Task `json:"result"`
}

type TaskRequest struct {
	Status string `json:"status"`
	Result Task `json:"result"`
}
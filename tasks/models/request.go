package models

type TasksRequest struct {
	Status string `json:"status"`
	Result []Task `json:"result"`
}
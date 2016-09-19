package models

type TasksRequest struct {
	Status string        `json:"status"`
	Result []interface{} `json:"result"`
}

type TaskRequest struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

package mrs_service_go_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tlantic/mrs-service-go-sdk/tasks/models"
)

const TASKS_MODULE_NAME = "tasks"

// Retrieves a task summary
func (c *Client) Summary() (*models.TaskSummary, error) {
	p := models.TaskSummary{}

	url := fmt.Sprintf("%s/%s", c.BaseApi, createPath(c.Organization, c.Application, "summary"))

	req, err := http.NewRequest("GET", url, nil)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

// GetTasks retrieves all available tasks.
//
// Whe items is true,  it also retrieves its items.
func (c *Client) GetTasks(items bool) (*models.TasksRequest, error) {
	p := models.TasksRequest{}

	url := fmt.Sprintf("%s/%s", c.BaseApi, createPath(c.Organization, c.Application, fmt.Sprintf("?items=%t", items)))

	req, err := http.NewRequest("GET", url, nil)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

// Create a task and its items.
func (c *Client) CreateTask(task *models.Task) (*models.TasksRequest, error) {
	p := models.TasksRequest{}

	url := fmt.Sprintf("%s/%s", c.BaseApi, createPath(c.Organization, c.Application, ""))

	req, err := http.NewRequest("POST", url, nil)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

func createPath(org, app, extra string) string {
	return fmt.Sprintf("%s/orgs/%s/apps/%s/%s", TASKS_MODULE_NAME, org, app, extra)
}

// GetAllTaskSetups : get all task setups
func (c *Client) GetAllTaskSetups() (*models.TaskRequest, error) {

	p := models.TaskRequest{}

	url := fmt.Sprintf("%s/%s/%s", c.BaseApi, createPath(c.Organization, c.Application, "task-setup"), "")

	req, err := http.NewRequest("GET", url, nil)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

// GetTaskSetup : get the task setup passed as argument
func (c *Client) GetTaskSetup(setup string) (*models.TaskRequest, error) {

	p := models.TaskRequest{}

	url := fmt.Sprintf("%s/%s/%s", c.BaseApi, createPath(c.Organization, c.Application, "task-setup"), setup)

	req, err := http.NewRequest("GET", url, nil)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil
}

// CreateTaskSetup : create a new task setup
func (c *Client) CreateTaskSetup(taskSetup *models.TaskSetup) (*models.TaskRequest, error) {

	p := models.TaskRequest{}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(taskSetup)

	url := fmt.Sprintf("%s/%s/%s", c.BaseApi, createPath(c.Organization, c.Application, "task-setup"), "")

	req, err := http.NewRequest("POST", url, b)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil

}

// UpdateTaskSetup : update "taskSetupId" with the info given in "taskSetup"
func (c *Client) UpdateTaskSetup(taskSetupId string, taskSetup interface{}) (*models.TaskRequest, error) {

	p := models.TaskRequest{}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(taskSetup)

	url := fmt.Sprintf("%s/%s/%s", c.BaseApi, createPath(c.Organization, c.Application, "task-setup"), taskSetupId)

	req, err := http.NewRequest("PUT", url, b)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	return &p, nil

}

// DeleteTaskSetup : delete "taskSetup"
func (c *Client) DeleteTaskSetup(taskSetup string) error {

	url := fmt.Sprintf("%s/%s/%s", c.BaseApi, createPath(c.Organization, c.Application, "task-setup"), taskSetup)

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	err = c.Send(req, nil)

	return err

}

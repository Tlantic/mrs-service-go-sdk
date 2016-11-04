package mrs_service_go_sdk


import (
"net/http"
"fmt"

"github.com/Tlantic/mrs-service-go-sdk/tasks/models"
)

const  TASKS_MODULE_NAME  = "tasks"

// Retrieves a task summary
func (c *Client) Summary() (*models.TaskSummary, error) {
	p := models.TaskSummary{}

	url := fmt.Sprintf("%s/%s", c.BaseApi, createPath(c.Organization, c.Application,"summary"))

	req, err := http.NewRequest("GET", url	, nil)
	req.Header.Set("mrs-application-id", c.AppId)

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

	req, err := http.NewRequest("GET", url	, nil)
	req.Header.Set("mrs-application-id", c.AppId)

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
func (c *Client) CreateTask( task *models.Task ) (*models.TasksRequest, error) {
	p := models.TasksRequest{}

	url := fmt.Sprintf("%s/%s", c.BaseApi, createPath(c.Organization, c.Application, ""))

	req, err := http.NewRequest("POST", url	, nil)
	req.Header.Set("mrs-application-id", c.AppId)

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}


	return &p, nil
}

func createPath(org, app, extra string) string{
	return fmt.Sprintf("%s/orgs/%s/apps/%s/%s/%s", TASKS_MODULE_NAME, org, app, TASKS_MODULE_NAME, extra)
}


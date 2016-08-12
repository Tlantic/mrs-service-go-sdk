package mrs_service_go_sdk


import (
	"net/http"
	"fmt"
)

const  TASKS_MODULE_NAME  = "tasks"

func (c *Client) Summary() (*TaskSummary, error) {
	p := TaskSummary{}

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


func (c *Client) GetTasks(items bool) (*TasksRequest, error) {
	p := TasksRequest{}

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

func createPath(org, app, extra string) string{
	return fmt.Sprintf("%s/orgs/%s/apps/%s/%s/%s", TASKS_MODULE_NAME, org, app, TASKS_MODULE_NAME, extra)
}


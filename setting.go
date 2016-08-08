package mrs_service_go_sdk

import (
	"net/http"
	"fmt"
)

func (c *Client) GetSetting(setting string) (*Setting, error) {
	p := Setting{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/%s", c.BaseApi, c.Organization, c.Application, "settings/"+setting), nil)
	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}


	return &p, nil
}

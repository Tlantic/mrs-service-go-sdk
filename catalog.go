package mrs_service_go_sdk

import (
	"net/http"
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/Tlantic/mrs-service-go-sdk/catalog/models"
)

const CATALOG_MODULE_NAME = "catalog"


// GetSnapshot retrieves a snapshot from the rp for a given product.
//
// Whe items is true,  it also retrieves its items.
func (c *Client) GetSnapshot(audits *models.SnapshotRequest, adapter string) (*models.SnapshotResponse, error) {
	p := models.SnapshotResponse{}

	url := fmt.Sprintf("%s/%s", c.BaseApi, createServicePath(c.Organization, c.Application, fmt.Sprintf("snapshot/sku?adapter=%s", adapter)))

	jsonStr, err := json.Marshal(audits)

	if err != nil {
		fmt.Println("Panic....")
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("mrs-application-id", c.AppId)

	if err != nil {
		return &p, err
	}

	err = c.Send(req, &p)
	if err != nil {
		return &p, err
	}

	fmt.Println(&p)
	return &p, nil
}

func createServicePath(org, app, extra string) string {
	return fmt.Sprintf("%s/orgs/%s/apps/%s/%s", CATALOG_MODULE_NAME, org, app, extra)
}



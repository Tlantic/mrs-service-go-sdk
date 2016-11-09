package mrs_service_go_sdk

import (
	"net/http"
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/Tlantic/mrs-service-go-sdk/catalog/models"
)

const CATALOG_MODULE_NAME = "catalog"

func (c *Client) GetProductSnapshot(product string, retailStoreId string, startDate string, endDate string, options map[string]string) (*models.SnapshotResponse, error) {
	p := models.SnapshotResponse{}

	adapter := options["adapter"]

	url := fmt.Sprintf("%s/%s", c.BaseApi, createServicePath(c.Organization, c.Application, fmt.Sprintf("snapshot/product?adapter=%s", adapter)))

	request := models.SnapshotRequest{
		Product:product,
		StoreId: retailStoreId,
		StartDate: startDate,
		EndDate: endDate,
	}

	jsonStr, err := json.Marshal(request)

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
	return &p, nil
}

func (c *Client) GetProductsSnapshot(products []string, retailStoreId string, startDate string, endDate string, options map[string]string) (*models.SnapshotsResponse, error) {
	p := models.SnapshotsResponse{}

	adapter := options["adapter"]
	url := fmt.Sprintf("%s/%s", c.BaseApi, createServicePath(c.Organization, c.Application, fmt.Sprintf("snapshot/products?adapter=%s", adapter)))

	request := models.SnapshotsRequest{
		Products:products,
		StoreId: retailStoreId,
		StartDate: startDate,
		EndDate: endDate,
	}

	jsonStr, err := json.Marshal(request)

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
	return &p, nil
}

func createServicePath(org, app, extra string) string {
	return fmt.Sprintf("%s/orgs/%s/apps/%s/%s", CATALOG_MODULE_NAME, org, app, extra)
}



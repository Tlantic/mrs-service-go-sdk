package mrs_service_go_sdk

import (
	"testing"
	"fmt"
	"github.com/Tlantic/mrs-service-go-sdk/catalog/models"
)

func TestGetSnapshot(t *testing.T) {

	cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "db22cd20-b668-4227-8a8d-92267f09af34")
	if err != nil {
		t.Errorf(err.Error())
	}

	req := models.SnapshotRequest{}

	req.Products = []string{"12754"}
	req.StoreId = "S0002"
	req.StartDate = "04-11-2016"
	req.EndDate = "04-11-2016"

	resp, err := cl.GetSnapshot(&req, "cockpit-adapter")

	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(resp)
	}

	t.Log(resp)
}
package mrs_service_go_sdk

import (
	"testing"
	"fmt"
	"time"
	"github.com/Tlantic/mrs-service-go-sdk/catalog/models"
)

func TestGetSnapshot(t *testing.T) {

	//cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "db22cd20-b668-4227-8a8d-92267f09af34")
	cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "")
	if err != nil {
		t.Errorf(err.Error())
	}

	startDate := time.Now().Format("02-01-2006")
	endDate := startDate

	//products := [2]string{"24649", "246972"}

	request := models.SnapshotRequest{
		StoreId: "S0001",
		StartDate: startDate,
		EndDate: endDate,
	}

	request.Products = append(request.Products, "24649")
	request.Products = append(request.Products, "246972")

	fmt.Println(startDate)

	resp, err := cl.GetProductsSnapshot(request, "cockpit-adapter")

	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(resp)
	}

	t.Log(resp)
}
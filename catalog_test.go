package mrs_service_go_sdk

import (
	"testing"
	"fmt"
	"time"
)

func TestGetProductsSnapshot(t *testing.T) {

	//cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "db22cd20-b668-4227-8a8d-92267f09af34")
	cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "")
	if err != nil {
		t.Errorf(err.Error())
	}

	options := map[string]string{
		"adapter": "cockpit-adapter",
	}

	products := []string{"24649", "246972"}
	storeId := "S0001"
	startDate := time.Now().Format("02-01-2006")
	endDate := startDate

	resp, err := cl.GetProductsSnapshot(products, storeId, startDate, endDate, options)

	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(resp)
	}

	t.Log(resp)
}

func TestGetProductSnapshot(t *testing.T) {

	//cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "db22cd20-b668-4227-8a8d-92267f09af34")
	cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "")
	if err != nil {
		t.Errorf(err.Error())
	}

	startDate := time.Now().Format("02-01-2006")
	endDate := startDate

	var options map[string]string = make(map[string]string)
	options["adapter"] = "cockpit-adapter"

	resp, err := cl.GetProductSnapshot("24649", "S0001", startDate, endDate, options)

	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(resp)
	}

	t.Log(resp)
}
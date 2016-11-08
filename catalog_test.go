package mrs_service_go_sdk

import (
	"testing"
	"fmt"
)

func TestGetSnapshot(t *testing.T) {

	//cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "db22cd20-b668-4227-8a8d-92267f09af34")
	cl, err := NewClient("http://52.50.91.27:9999", "monteserrat", "cockpit", "")
	if err != nil {
		t.Errorf(err.Error())
	}

	resp, err := cl.GetSnapshot("12754", "S0002", "08-11-2016", "08-11-2016", "cockpit-adapter")

	if err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(resp)
	}

	t.Log(resp)
}
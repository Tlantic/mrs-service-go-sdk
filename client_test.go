package mrs_service_go_sdk

import (
	"fmt"
	"testing"
)


func TestNewClient(t *testing.T) {
	_, err := NewClient("", "", "")
	if err == nil {
		t.Errorf("All arguments are required in NewClient()")
	} else {
		fmt.Println(err.Error())
	}

	_, err = NewClient("http://52.50.91.:9999","tlantic", "customer")
	if err != nil {
		t.Errorf( err.Error())
	}


}

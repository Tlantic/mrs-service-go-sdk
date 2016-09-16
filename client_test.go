package mrs_service_go_sdk

import (
	"fmt"
	"testing"
)


func TestNewClient(t *testing.T) {
	_, err := NewClient("", "", "", map[string]string{})
	if err == nil {
		t.Errorf("All arguments are required in NewClient()")
	} else {
		fmt.Println(err.Error())
	}

	_, err = NewClient("http://52.50.91.27","tlantic", "customer",map[string]string{})
	if err != nil {
		t.Errorf( err.Error())
	}


}

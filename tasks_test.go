package mrs_service_go_sdk

import (
	"testing"
)


func TestTaskSummary(t *testing.T) {

	cl, err := NewClient("http://52.50.91.27:8067","tlantic", "instore", "489a598e-259c-4e25-974e-5de00b29f707")
	if err != nil {
		t.Errorf( err.Error())
	}

	summary, err := cl.Summary()

	if  err != nil{
		t.Errorf( err.Error())
	}

	t.Log(summary.Result.PriceChang.Total)


}


func TestGetTasks(t *testing.T) {

	cl, err := NewClient("http://52.50.91.27:8067","tlantic", "instore", "489a598e-259c-4e25-974e-5de00b29f707")
	if err != nil {
		t.Errorf( err.Error())
	}

	tasks, err := cl.GetTasks(true)

	if  err != nil{
		t.Errorf( err.Error())
	}

	t.Log(tasks.Result)

	if(len(tasks.Result)>0){
		task := tasks.Result[0]
		t.Log(task.TaskName)
		t.Log(task.Items)
	}


}


package mrs_service_go_sdk

import (
	"testing"
	"github.com/Tlantic/mrs-service-go-sdk/tasks/models"
)

func TestTaskSummary(t *testing.T) {

	cl, err := NewClient("http://52.50.91.27:8067", "tlantic", "instore", "489a598e-259c-4e25-974e-5de00b29f707")
	if err != nil {
		t.Errorf(err.Error())
	}

	summary, err := cl.Summary()

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(summary.Result.PriceChang.Total)

}

func TestGetTasks(t *testing.T) {

	cl, err := NewClient("http://52.50.91.27:8067", "tlantic", "instore", "489a598e-259c-4e25-974e-5de00b29f707")
	if err != nil {
		t.Errorf(err.Error())
	}

	tasks, err := cl.GetTasks(true)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(tasks.Result)

	if (len(tasks.Result) > 0) {
		task := tasks.Result[0]
		t.Log(task.TaskName)
		t.Log(task.Items)
	}

}

func TestGetTaskById(t *testing.T) {

	cl, err := NewClient("http://52.50.91.27:8067", "tlantic", "instore", "489a598e-259c-4e25-974e-5de00b29f707")
	if err != nil {
		t.Errorf(err.Error())
	}

	tasks, err := cl.GetTaskById("589f3517-0c98-4071-9fea-ab3e91dbe7ec", true)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(tasks.Result)

	task := tasks.Result
	t.Log(task.TaskName)
	t.Log(task.Items)

	for _, item := range task.Items {
		t.Log(models.ProductTaskItem(item))
	}

	t.Log(task.Items)
}

func TestCreateTask(t *testing.T) {

	cl, err := NewClient("http://52.50.91.27", "tlantic", "instore", "")
	if err != nil {
		t.Errorf(err.Error())
	}

	tasks, err := cl.CreateTask(&models.Task{

	})

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(tasks.Result)

	if (len(tasks.Result) > 0) {
		task := tasks.Result[0]
		t.Log(task.TaskName)
		t.Log(task.Items)
	}

}

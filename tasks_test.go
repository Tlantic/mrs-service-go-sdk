package mrs_service_go_sdk

import (
	"testing"

	"github.com/Tlantic/mrs-service-go-sdk/tasks/models"
)

func TestTaskSummary(t *testing.T) {

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://52.50.91.27:8067", "tlantic", "instore", headers)
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

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://52.50.91.27:8067", "tlantic", "instore", headers)
	if err != nil {
		t.Errorf(err.Error())
	}

	tasks, err := cl.GetTasks(true)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(tasks.Result)

}

func TestCreateTask(t *testing.T) {

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://52.50.91.27:8067", "tlantic", "instore", headers)
	if err != nil {
		t.Errorf(err.Error())
	}

	tasks, err := cl.CreateTask(&models.Task{})

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(tasks.Result)

}

// TestGetAllTaskSetups
func TestGetAllTaskSetups(t *testing.T) {

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://localhost:8010", "tlantic", "instore", headers)
	if err != nil {
		t.Errorf(err.Error())
	}

	taskRequest, err := cl.GetAllTaskSetups()

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(taskRequest.Result)

}

// TestGetSetup
func TestGetTaskSetup(t *testing.T) {

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://localhost:8010", "tlantic", "instore", headers)
	if err != nil {
		t.Errorf(err.Error())
	}

	taskRequest, err := cl.GetTaskSetup("Setup1")

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(taskRequest.Result)

}

// TestCreateTaskSetup ...
func TestCreateTaskSetup(t *testing.T) {

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://localhost:8010", "tlantic", "instore", headers)
	if err != nil {
		t.Errorf(err.Error())
	}

	taskSetupName := models.TaskSetupName{
		Value:      "TaskName1",
		LanguageId: "pt-PT",
	}

	taskLangs := []models.TaskSetupName{taskSetupName}
	taskAttributes := models.TaskSetupAttributes{}

	taskRequest, err := cl.CreateTaskSetup(&models.TaskSetup{
		SetupId:         "Setup5",
		Status:          "B",
		TaskType:        "TaskTypeGo",
		TaskName:        taskLangs,
		TaskDescription: "TaskDescriptionGO",
		Periodicity:     "Monthly",
		LastRunTime:     123123123,
		GroupIds:        []string{"GO1", "GO2", "GO3", "GO4", "GO5", "GO6"},
		StoreIds:        []string{"Store1", "Store2", "Store3", "Store4", "Store5"},
		ApplicationIds:  []string{"instore", "customer"},
		Attributes:      taskAttributes,
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(taskRequest.Result)

}

// TestUpdateTaskSetup ...
func TestUpdateTaskSetup(t *testing.T) {

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://localhost:8010", "tlantic", "instore", headers)
	if err != nil {
		t.Errorf(err.Error())
	}

	taskLangs := []models.TaskSetupName{
		models.TaskSetupName{
			Value:      "TaskName1_changed_to_en",
			LanguageId: "en-EN",
		},
		models.TaskSetupName{
			Value:      "TaskName1_changed_to_fr",
			LanguageId: "fr-FR",
		},
	}

	taskRequest, err := cl.UpdateTaskSetup("Setup1", &models.TaskSetup{
		TaskName: taskLangs,
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(taskRequest.Result)
}

// TestDeleteTaskSetup ...
func TestDeleteTaskSetup(t *testing.T) {

	headers := map[string]string{
		"mrs-application-id":  "489a598e-259c-4e25-974e-5de00b29f707",
		"mrs-organization-id": "1ecf2c6d-27d3-4bb0-b319-cd51c2797c1d",
	}

	cl, err := NewClient("http://localhost:8010", "tlantic", "instore", headers)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = cl.DeleteTaskSetup("Setup1")

	if err != nil {
		t.Errorf(err.Error())
	}

}

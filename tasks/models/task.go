package models

type Task struct {
	UID                string                                `json:"_uId"`
	ParentTaskId       string                                `json:"parentTaskId"`
	AlertId            string                                `json:"alertId"`
	TaskName           string                                `json:"taskName"`
	TaskType           string                                `json:"taskType"`
	TaskSubType        string                                `json:"taskSubType"`
	Items              []TaskItem                        `json:"items"`
	Description        string                                `json:"description"`
	ProfileID          string                                `json:"profileId"`
	DeviceID           string                                `json:"deviceId"`
	Status             string                                `json:"status"`
	OwnerID            string                                `json:"ownerId"`
	LastUpdateUser     string                                `json:"lastUpdateUser"`
	Counter            int                                        `json:"counter"`
	SetupID            string                                `json:"setupId"`
	StoreID            string                                `json:"storeId"`
	Priority           int                                `json:"priority"`
	Shared             bool                                `json:"shared"`
	DocumentID         string                                `json:"documentId"`
	ScheduledDate      int64                                `json:"scheduledDate"`
	ExpectedFinishDate int64                                `json:"expectedFinishDate"`
	StartDate          int64                                `json:"startDate"`
	FinishDate         int64                                `json:"finishDate"`
	ItemGroupID        string                                `json:"itemGroupID"`

	ApplicationID      string                                `json:"applicationId"`
	Type               string                                `json:"_type"`
	UpdateDate         int                                        `json:"_updateDate"`
	CreateDate         int                                `json:"_createDate"`
	UpdateUser         string                                `json:"_updateUser"`
	CreateUser         string                                `json:"_createUser"`
	ItemType           interface{}                     `json:"itemType"`
}
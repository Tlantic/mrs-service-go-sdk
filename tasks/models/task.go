package models

import "time"

type Task struct {

	Id 							string 					`json:"_uId" validate:"nonzero"`
	Type 						string 					`json:"_type"`
	Status						string					`json:"status" validate:"nonzero, min=1, max=2"`

	// Meta Data
	ApplicationID 				string 					`json:"applicationId" validate:"nonzero"`
	SetupId 					string 					`json:"setupId,omitempty" validate:"nonzero"`
	ParentTaskId 				string 					`json:"parentTaskId,omitempty" validate:"nonzero"`
	StoreId 					string 					`json:"storeId" validate:"nonzero"`
	ProcessId 					string 					`json:"processId,omitempty" validate:"nonzero"`
	ProfileId 					string 					`json:"profileId,omitempty" validate:"nonzero"`
	DeviceId 					string 					`json:"deviceId,omitempty" validate:"nonzero"`
	OwnerId 					string 					`json:"ownerId"`
	UpdateDate 					uint					`json:"_updateDate" validate:"nonzero"`
	CreateDate 					uint 					`json:"_createDate" validate:"nonzero"`
	UpdateUser 					string 					`json:"_updateUser"`
	CreateUser 					string 					`json:"_createUser"`


	// Description Data
	TaskType					string					`json:"taskType" validate:"nonzero"`
	TaskName					string					`json:"taskName" validate:"nonzero"`
	TaskName_gb					string					`json:"taskName_gb,omitempty" validate:"nonzero"`
	TaskName_pt					string					`json:"taskName_es,omitempty" validate:"nonzero"`
	TaskName_es					string					`json:"taskName_es,omitempty" validate:"nonzero"`
	TaskName_de					string					`json:"taskName_de,omitempty" validate:"nonzero"`
	TaskName_us					string					`json:"taskName_us,omitempty" validate:"nonzero"`
	TaskName_br					string					`json:"taskName_br,omitempty" validate:"nonzero"`
	TaskName_ve					string					`json:"taskName_ve,omitempty" validate:"nonzero"`

	Description					string					`json:"description" validate:"nonzero"`
	Description_gb				string					`json:"description_gb,omitempty" validate:"nonzero"`
	Description_pt				string					`json:"description_es,omitempty" validate:"nonzero"`
	Description_es				string					`json:"description_es,omitempty" validate:"nonzero"`
	Description_de				string					`json:"description_de,omitempty" validate:"nonzero"`
	Description_us				string					`json:"description_us,omitempty" validate:"nonzero"`
	Description_br				string					`json:"description_br,omitempty" validate:"nonzero"`
	Description_ve				string					`json:"description_ve,omitempty" validate:"nonzero"`

	// Operational Data
	Priority					uint					`json:"priority"`
	ItemLimit					uint					`json:"itemLimit"`
	IsReleasable				bool					`json:"isReleasable"`
	IsMergeable					bool					`json:"isMergeable"`
	IsOriented					bool					`json:"isOriented"`

	// Variable Data
	Attributes					map[string]interface{}	`json:"attributes"`

	// Scheduling Data
	ScheduledDate				time.Time				`json:"scheduledDate" validate:"nonzero"`
	ExpectedFinishDate			time.Time				`json:"expectedFinishDate,omitempty" validate:"nonzero"`
	StartDate					time.Time				`json:"expectedFinishDate"`
	StartUser					string					`json:"startUser"`
	ReopenDate					time.Time				`json:"reopenDate"`
	ReopenUser					string					`json:"reopenUser"`
	FinishDate					time.Time				`json:"finishDate"`
	FinishUser					string					`json:"finishUser"`

}
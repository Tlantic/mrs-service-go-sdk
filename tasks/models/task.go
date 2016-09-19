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

	Name						string					`json:"name"`
	Name_gb						string					`json:"name_gb,omitempty" validate:"nonzero"`
	Name_pt						string					`json:"name_es,omitempty" validate:"nonzero"`
	Name_es						string					`json:"name_es,omitempty" validate:"nonzero"`
	Name_de						string					`json:"name_de,omitempty" validate:"nonzero"`
	Name_us						string					`json:"name_us,omitempty" validate:"nonzero"`
	Name_br						string					`json:"name_br,omitempty" validate:"nonzero"`
	Name_ve						string					`json:"name_ve,omitempty" validate:"nonzero"`

	Description					string					`json:"description"`
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
	IsAcceptingItems			bool					`json:"isAcceptingItems"`

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
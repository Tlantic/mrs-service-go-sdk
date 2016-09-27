package models

import "time"


type Task_Validations struct {
	 Errors						map[string]string		`json:"errors"`
	 Warnings 					map[string]string		`json:"warnings"`
}

type Task struct {

	// Meta
	Id 							string 					`json:"_uId" validate:"nonzero"`
	Type 						string 					`json:"_type"`
	CreateDate 					time.Time 				`json:"_createDate" validate:"nonzero"`
	UpdateDate 					time.Time				`json:"_updateDate" validate:"nonzero"`
	CreateUser 					string 					`json:"_createUser"`
	UpdateUser 					string 					`json:"_updateUser"`
	Status						string					`json:"status" validate:"nonzero, min=1, max=2"`
	ApplicationId 				string 					`json:"applicationId" validate:"nonzero"`
	OrganizationId 				string 					`json:"organizationId" validate:"nonzero"`

	// (cont.) Meta Data
	SetupId 					string 					`json:"setupId,omitempty" validate:"nonzero"`
	ParentTaskId 				string 					`json:"parentTaskId,omitempty" validate:"nonzero"`
	StoreId 					string 					`json:"storeId" validate:"nonzero"`
	ProcessId 					string 					`json:"processId,omitempty" validate:"nonzero"`
	ProfileId 					string 					`json:"profileId,omitempty" validate:"nonzero"`
	DeviceId 					string 					`json:"deviceId,omitempty" validate:"nonzero"`
	OwnerId 					string 					`json:"ownerId"`





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
	IsAcceptingItems			bool					`json:"isAcceptingItems"`
	IsReleasable				bool					`json:"isReleasable"`
	IsReasignable				bool					`json:"isReassignable"`
	IsOriented					bool					`json:"isOriented"`


	// Variable Data
	Attributes					map[string]interface{}	`json:"attributes"`
	Validations					Task_Validations		`json:"validations"`

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
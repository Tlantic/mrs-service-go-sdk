package models

import "time"


type TaskSetup struct {

	// Meta
	Id 							string 					`json:"_uId" validate:"nonzero"`
	Type 						string 					`json:"_type"`
	CreateDate 					time.Time 				`json:"_createDate" validate:"nonzero"`
	UpdateDate 					time.Time				`json:"_updateDate" validate:"nonzero"`
	CreateUser 					string 					`json:"_createUser"`
	UpdateUser 					string 					`json:"_updateUser"`
	Status						string					`json:"status" validate:"nonzero, min=1, max=2"`
	OrganizationId  			string              	`json:"organizationId" validate:"nonzero"`

	// (cont.) Meta Data
	SetupId             		string              	`json:"setupId" validate:"nonzero"`
	StoreIds        			[]string            	`json:"storeIds"`
	ApplicationIds  			[]string            	`json:"applicationIds"`
	GroupIds  					[]string            	`json:"groupIds"`


	// Task Descriptive
	TaskType					string					`json:"taskType" validate:"nonzero"`

	TaskName					string					`json:"name"`
	TaskName_gb					string					`json:"name_gb,omitempty" validate:"nonzero"`
	TaskName_pt					string					`json:"name_es,omitempty" validate:"nonzero"`
	TaskName_es					string					`json:"name_es,omitempty" validate:"nonzero"`
	TaskName_de					string					`json:"name_de,omitempty" validate:"nonzero"`
	TaskName_us					string					`json:"name_us,omitempty" validate:"nonzero"`
	TaskName_br					string					`json:"name_br,omitempty" validate:"nonzero"`
	TaskName_ve					string					`json:"name_ve,omitempty" validate:"nonzero"`

	TaskDescription				string					`json:"description"`
	TaskDescription_gb			string					`json:"description_gb,omitempty" validate:"nonzero"`
	TaskDescription_pt			string					`json:"description_es,omitempty" validate:"nonzero"`
	TaskDescription_es			string					`json:"description_es,omitempty" validate:"nonzero"`
	TaskDescription_de			string					`json:"description_de,omitempty" validate:"nonzero"`
	TaskDescription_us			string					`json:"description_us,omitempty" validate:"nonzero"`
	TaskDescription_br			string					`json:"description_br,omitempty" validate:"nonzero"`
	TaskDescription_ve			string					`json:"description_ve,omitempty" validate:"nonzero"`

	Periodicity    				string		           	`json:"periodicity"`
	Attributes      			map[string]interface{} 	`json:"attributes"`
}

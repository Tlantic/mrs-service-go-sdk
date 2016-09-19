package models

type TaskItem struct {

	Id 							string 					`json:"_uId" validate:"nonzero"`
	Type 						string 					`json:"_type"`
	Status						string					`json:"status" validate:"nonzero, min=1, max=2"`


	// Meta Data
	ApplicationID 				string 					`json:"applicationId" validate:"nonzero"`
	TaskId		 				string 					`json:"taskId" validate:"nonzero"`

	UpdateDate 					uint					`json:"_updateDate" validate:"nonzero"`
	CreateDate 					uint 					`json:"_createDate" validate:"nonzero"`
	UpdateUser 					string 					`json:"_updateUser"`
	CreateUser 					string 					`json:"_createUser"`


	// Description Data
	Name						string					`json:"name" validate:"nonzero"`
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
	ItemId                		string 					`json:"itemId"`
	StartDate             		int 					`json:"startDate"`

	Attributes				map[string]interface{}		`json:"attributes"`
}
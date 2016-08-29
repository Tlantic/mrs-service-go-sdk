package models

type ChecklistTaskItem_v1 struct {

	UniqueId			  	string						`json:"_uId"`
	ApplicationId 			string						`json:"applicationId"`
	TaskId					string						`json:"taskId"`
	Type		  			string						`json:"type"`
	Status                	string				        `json:"status"`
	Observation           	string 						`json:"observation"`
	StartDate             	int 						`json:"startDate"`


	ItemId                	string 						`json:"itemId"`
}


func ( item *ChecklistTaskItem_v1  ) GetUniqueId() string {
	return item.UniqueId
}

func ( item *ChecklistTaskItem_v1  ) GetItemId() string {
	return item.ItemId
}

func ( item *ChecklistTaskItem_v1  ) GetType() string {
	return item.Type
}

func ( item *ChecklistTaskItem_v1  ) GetStatus() string {
	return item.Status
}

func ( item *ChecklistTaskItem_v1  ) GetStartDate() int {
	return item.StartDate
}

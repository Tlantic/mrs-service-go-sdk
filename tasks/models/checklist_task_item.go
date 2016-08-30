package models

type ChecklistTaskItem struct {

	UniqueId			  	string						`json:"_uId"`
	ApplicationId 			string						`json:"applicationId"`
	TaskId					string						`json:"taskId"`
	Type		  			string						`json:"type"`
	Status                	string				        `json:"status"`
	Observation           	string 						`json:"observation"`
	StartDate             	int 						`json:"startDate"`


	ItemId                	string 						`json:"itemId"`
}


func ( item *ChecklistTaskItem  ) GetUniqueId() string {
	return item.UniqueId
}

func ( item *ChecklistTaskItem  ) GetItemId() string {
	return item.ItemId
}

func ( item *ChecklistTaskItem  ) GetType() string {
	return item.Type
}

func ( item *ChecklistTaskItem  ) GetStatus() string {
	return item.Status
}

func ( item *ChecklistTaskItem  ) GetStartDate() int {
	return item.StartDate
}

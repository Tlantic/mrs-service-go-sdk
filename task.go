package mrs_service_go_sdk

type TaskSummary struct {
	Status string `json:"status"`
	Result struct {
		       PriceAudit struct {
					  Total int `json:"total"`
				  } `json:"priceAudit"`
		       StockCount struct {
					  Total int `json:"total"`
				  } `json:"stockCount"`
		       StockOut struct {
					  Total int `json:"total"`
				  } `json:"stockOut"`
		       Replenish struct {
					  Total int `json:"total"`
				  } `json:"replenish"`
		       QueueBust struct {
					  Total int `json:"total"`
				  } `json:"queueBust"`
		       Checklist struct {
					  Total int `json:"total"`
				  } `json:"checklist"`
		       Shrinkage struct {
					  Total int `json:"total"`
				  } `json:"shrinkage"`
		       Planogram struct {
					  Total int `json:"total"`
				  } `json:"planogram"`
		       Markdown struct {
					  Total int `json:"total"`
				  } `json:"markdown"`
		       QualClas struct {
					  Total int `json:"total"`
				  } `json:"qualClas"`
		       RecepDir struct {
					  Total int `json:"total"`
				  } `json:"recepDir"`
		       RecepCen struct {
					  Total int `json:"total"`
				  } `json:"recepCen"`
		       RecepTsf struct {
					  Total int `json:"total"`
				  } `json:"recepTsf"`
		       PriceChang struct {
					  Total int `json:"total"`
				  } `json:"priceChang"`
		       RecepCS struct {
					  Total int `json:"total"`
				  } `json:"recepCS"`
		       ReplenPrep struct {
					  Total int `json:"total"`
				  } `json:"replenPrep"`
		       ReplenConf struct {
					  Total int `json:"total"`
				  } `json:"replenConf"`
		       StockOutFA struct {
					  Total int `json:"total"`
				  } `json:"stockOutFA"`
	       } `json:"result"`
}

type TasksRequest struct {
	Status string `json:"status"`
	Result []Task `json:"result"`
}


type Task struct {
	Status string `json:"status"`
	OwnerID string `json:"ownerId"`
	LastUpdateUser string `json:"lastUpdateUser"`
	TaskType string `json:"taskType"`
	Counter int `json:"counter"`
	SetupID string `json:"setupId"`
	TaskName string `json:"taskName"`
	Description string `json:"description"`
	ProfileID string `json:"profileId"`
	DeviceID string `json:"deviceId"`
	StoreID string `json:"storeId"`
	Priority int `json:"priority"`
	Shared bool `json:"shared"`
	DocumentID string `json:"documentId"`
	ScheduledDate int64 `json:"scheduledDate"`
	ApplicationID string `json:"applicationId"`
	Type string `json:"_type"`
	UID string `json:"_uId"`
	UpdateDate int `json:"_updateDate"`
	CreateDate int `json:"_createDate"`
	UpdateUser string `json:"_updateUser"`
	CreateUser string `json:"_createUser"`
	Items []TaskItem `json:"items"`
}



type TaskItem struct {
	ItemID string `json:"itemId"`
	Status string `json:"status"`
	TaskID string `json:"taskId"`
	Type string `json:"type,omitempty"`
	ApplicationID string `json:"applicationId"`
	_Type string `json:"_type"`
	UID string `json:"_uId"`
	UpdateDate int `json:"_updateDate"`
	CreateDate int `json:"_createDate"`
	UpdateUser string `json:"_updateUser,omitempty"`
	CreateUser string `json:"_createUser,omitempty"`
	ItemEan string `json:"itemEan,omitempty"`
}
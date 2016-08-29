package models


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
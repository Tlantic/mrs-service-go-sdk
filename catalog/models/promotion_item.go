package models

import "time"

type PromotionItem struct{
	Type					string						`json:"type"`
	GroupId 				string   					`json:"groupId"`
	StoreId	 				string  					`json:"storeId"`
	StartDate 				time.Time 					`json:"startDate"`
	EndDate 				time.Time 					`json:"endDate"`
	Desc 					string 						`json:"desc"`
	Desc_pt 				string 						`json:"desc_pt"`
	Desc_en 				string 						`json:"desc_en"`
	Desc_br 				string 						`json:"desc_br"`
	Desc_es 				string 						`json:"desc_es"`
	Desc2 					string 						`json:"desc2"`
	Desc2_pt 				string 						`json:"desc2_pt"`
	Desc2_en 				string 						`json:"desc2_en"`
	Desc2_br 				string 						`json:"desc2_br"`
	Desc2_es 				string 						`json:"desc2_es"`
	Desc3 					string 						`json:"desc3"`
	Desc3_pt 				string 						`json:"desc3_pt"`
	Desc3_en 				string 						`json:"desc3_en"`
	Desc3_br 				string 						`json:"desc3_br"`
	Desc3_es 				string 						`json:"desc3_es"`
	Price 					float64 					`json:"price"`
}
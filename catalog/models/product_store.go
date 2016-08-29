package models

type ProductStore struct{
	StoreId					string 						`json:"storeId"`
	OldPrice				float64						`json:"oldPrice"`
	SellPrice				float64						`json:"sellPrice"`
	Discontinued			bool						`json:"discontinued"`
	Attributes				[]ProductAttribute	 		`json:"attributes"`
	Status                  string						`json:"status"`
	StockOnHand             float64						`json:"stockOnHand"`
}
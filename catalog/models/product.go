package models


type Product struct {

	Id 					string 							`json:"_id"`
	Name	 			string	 						`json:"name"`
	Name_pt 			string 							`json:"name_pt"`
	Name_en 			string 							`json:"name_en"`
	Name_br 			string 							`json:"name_br"`
	Name_es 			string 							`json:"name_es"`
	Desc 				string 							`json:"desc"`
	Desc_pt 			string 							`json:"desc_pt"`
	Desc_en 			string 							`json:"desc_en"`
	Desc_br 			string 							`json:"desc_br"`
	Desc_es 			string 							`json:"desc_es"`
	LongDesc 			string 							`json:"longDesc"`
	LongDesc_pt 		string 							`json:"longDesc_pt"`
	LongDesc_en 		string 							`json:"longDesc_en"`
	LongDesc_br 		string 							`json:"longDesc_br"`
	LongDesc_es 		string 							`json:"longDesc_es"`
	Status 				string 							`json:"status"`
	SupplierId 			string 							`json:"supplierId"`
	Brand 				string 							`json:"brand"`
	Brand_pt 			string 							`json:"brand_pt"`
	Brand_en 			string 							`json:"brand_en"`
	Brand_br 			string 							`json:"brand_br"`
	Brand_es 			string 							`json:"brand_es"`
	CategoryId 			string							`json:"categoryId"`
	MediaId				[]string 						`json:"mediaId"`
	HierarchyCategory	HierarchyCategory				`json:"hierarchyCategory"`
	Attributes			[]ProductAttribute 				`json:"attributes"`
	Stores 				[]ProductStore					`json:"stores"`
	ScanCodes			[]ScanCode						`json:"scanCodes"`
	Promotions 			[]PromotionItem					`json:"promotions"`
	Groups				[]GroupItem						`json:"groups"`
}
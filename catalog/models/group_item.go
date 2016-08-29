package models

type GroupItem struct{
	GroupId 			string  			`json:"groupId"`
	StoreId 			string  			`json:"storeId"`
	Products			[]string			`json:"products"`
}


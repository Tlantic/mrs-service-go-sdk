package models

type ScanCode struct{
	Ean					string 				`json:"ean"`
	Primary 			string				`json:"primary"`
	Type 				string				`json:"type"`
}
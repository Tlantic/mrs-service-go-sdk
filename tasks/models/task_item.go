package models

import "github.com/Tlantic/mrs-service-go-sdk/catalog/models"

type TaskItem struct {
	UniqueId              string                                                `json:"_uId"`
	ApplicationId         string                                                `json:"applicationId"`
	TaskId                string                                                `json:"taskId"`
	Type                  string                                                `json:"type"`
	Status                string                                        `json:"status"`
	Observation           string                                                `json:"observation"`
	StartDate             int                                                `json:"startDate"`

	ItemId                string                                                `json:"itemId"`
	ItemEAN               string                                                `json:"itemEAN"`
	ItemName              string                                                `json:"itemName"`
	ExpirationDate        int                                                `json:"expirationDate"`
	HierarchyCategory     models.HierarchyCategory        `json:"hierarchyCategory"`

	PosPrice              int                                                `json:"posPrice"`
	LabelPrice            int                                                `json:"labelPrice"`
	OldERPPrice           int                                                `json:"oldERPPrice"`
	PriceDivergence       bool                                                `json:"priceDivergence"`

	Quantity              int                                                `json:"quantity"`
	ExpectedQuantity      int                                                `json:"expectedQuantity"`

	CheckExpirationDate   bool                                                `json:"checkExpirationDate"`

	PickCount             int                                                `json:"pickCount"`
	FirstPickingTime      int                                                `json:"firstPickingTime"`
	ReceptionUnit         int                                                `json:"receptionUnit"`
	ReceptionUnitQuantity int                                                `json:"receptionUnitQuantity"`
}

func ( item *TaskItem  ) GetUniqueId() string {
	return item.UniqueId
}

func ( item *TaskItem  ) GetItemId() string {
	return item.ItemId
}

func ( item *TaskItem  ) GetItemEAN() string {
	return item.ItemEAN
}

func ( item *TaskItem  ) GetName() string {
	return item.ItemEAN
}

func ( item *TaskItem  ) GetType() string {
	return item.Type
}

func ( item *TaskItem  ) GetStatus() string {
	return item.Status
}

func ( item *TaskItem  ) GetStartDate() int {
	return item.StartDate
}


package models

type SnapshotRequest struct {
	Products  []string               `json:"products"`
	StoreId   string                 `json:"storeId"`
	StartDate string                 `json:"startDate"`
	EndDate   string                 `json:"endDate"`
}

type SnapshotResponse struct {
	Status string `json:"status"`
	Result map[string]Snapshot `json:"result"`
}

type Snapshot struct {
	Sku           int32              `json:"sku"`
	AvgCostPrice  float64              `json:"avgCostPrice"`
	PurchasePrice float64              `json:"purchasePrice"`
	PosPrice      float64              `json:"posPrice"`
	Stock         float64              `json:"stock"`
}


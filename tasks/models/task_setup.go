package models

type TaskSetupAttributes map[string]interface{}

type TaskSetupName struct {
	Value      string `json:"value"`
	LanguageId string `json:"languageId"`
}

type TaskSetup struct {
	Uid             string              `json:"_uId"`
	SetupId         string              `json:"setupId,omitempty"`
	OrganizationID  string              `json:"organizationId,omitempty"`
	Status          string              `json:"status,omitempty"`
	TaskType        string              `json:"taskType,omitempty"`
	TaskName        []TaskSetupName     `json:"taskName,omitempty"`
	TaskDescription string              `json:"taskDescription,omitempty"`
	Periodicity     string              `json:"periodicity,omitempty"`
	LastRunTime     int                 `json:"lastRunTime,omitempty"`
	GroupIds        []string            `json:"groupIds,omitempty"`
	StoreIds        []string            `json:"storeIds,omitempty"`
	ApplicationIds  []string            `json:"applicationIds,omitempty"`
	Attributes      TaskSetupAttributes `json:"attributes,omitempty"`
}

package models

type TaskItem interface {
	GetUniqueId()				string
	GetItemId()					string
	GetName()					string
	GetType()					string
	GetStatus()					string
	GetStartDate()				int
}
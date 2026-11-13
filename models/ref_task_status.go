package models

import "time"

type RefTaskStatus struct {
	ID         int    `gorm:"primary_key"`
	UUID       string `gorm:"unique"`
	TaskStatus string
	CreatedBy  string
	CreatedAt  time.Time
	UpdatedBy  string
	UpdatedAt  time.Time
	DeletedBy  string
	DeletedAt  time.Time
}

func (RefTaskStatus) TableName() string {
	return "ref_task_status"
}

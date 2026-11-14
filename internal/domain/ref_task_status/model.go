package ref_task_status

import "time"

type RefTaskStatusModels struct {
	ID         int        `gorm:"primary_key" json:"id,omitempty"`
	UUID       string     `gorm:"unique" json:"uuid,omitempty"`
	TaskStatus string     `json:"task_status,omitempty"`
	CreatedBy  string     `json:"created_by,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedBy  string     `json:"updated_by,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedBy  string     `json:"deleted_by,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func (RefTaskStatusModels) TableName() string {
	return "ref_task_status"
}

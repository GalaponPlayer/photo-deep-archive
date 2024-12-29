package entity

import "time"

type CommonDBAttributes struct {
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt  time.Time `json:"deleted_at"`
}

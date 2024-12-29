package entity

import "time"

type CommonDBAttributes struct {
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt time.Time `json:"modified_at" gorm:"autoUpdateTime"`
	DeletedAt  time.Time `json:"deleted_at"`
}

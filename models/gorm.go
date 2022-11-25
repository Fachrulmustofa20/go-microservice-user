package models

import "time"

type Gorm struct {
	ID        uint64     `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

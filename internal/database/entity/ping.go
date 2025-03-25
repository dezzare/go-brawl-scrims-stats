package entity

import "gorm.io/gorm"

type Ping struct {
	gorm.Model
	ID        uint
	TimeInMs  float64
	CreatedAt string
	UpdatedAt string
}

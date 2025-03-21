package entity

import "gorm.io/gorm"

type Ping struct {
	gorm.Model
	ID        string
	Host      string
	IP        string
	Time      float64
	CreatedAt string
	UpdatedAt string
}

package model

import "gorm.io/gorm"

type Ping struct {
	gorm.Model
	TimeInMs float64
}

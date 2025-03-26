package entity

import "gorm.io/gorm"

type Mode struct {
	gorm.Model
	Name string
	Maps []Map
}

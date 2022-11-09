package customermodel

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Ballance float64
}

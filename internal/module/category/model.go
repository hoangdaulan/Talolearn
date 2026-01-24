package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"size:100;not null" json:"name"`
}

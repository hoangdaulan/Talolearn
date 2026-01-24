package course

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string `gorm:"size:100;not null" json:"title"`
	Description string `gorm:"size:255" json:"description"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

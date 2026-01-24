package teacher

import (
	"Talolearn-Project/internal/module/category"
	"Talolearn-Project/internal/module/user"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	UserID     uint
	Bio        string `json:"bio"`
	CategoryID uint   `json:"category_id" binding:"required"`

	User     user.User         `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category category.Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TeacherRequest struct {
	Bio        string `json:"bio"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

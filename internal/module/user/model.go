package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `gorm:"size:100;not null" json:"name"`
	PasswordHash string `gorm:"size:255;not null" json:"password"`
	Email        string `gorm:"size:100;unique;not null" json:"email"`
	Avatar       string `gorm:"size:255" json:"avatar"`
	Phone        string `gorm:"phone;unique" json:"phone" binding:"required"`
	RoleID       uint   `gorm:"default:1"json:"role_id"`
}

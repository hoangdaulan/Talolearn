package student

import "Talolearn-Project/internal/module/user"

type Student struct {
	Code   string `json:"code" gorm:"primaryKey;size:20;not null;unique"`
	UserID uint   `json:"user_id" gorm:"not null;unique"`

	User user.User `gorm:"foreignKey:UserID;references:ID"`
}

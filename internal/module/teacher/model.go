package teacher

type Teacher struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:100;not null" json:"name"`
	Email  string `gorm:"size:100;unique;not null" json:"email"`
	Avatar string `gorm:"size:255" json:"avatar"`
	Phone  string `gorm:"phone;unique" json:"phone" binding:"required"`
}

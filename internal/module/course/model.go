package course

type Course struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"size:100;not null" json:"title"`
	Description string `gorm:"size:255" json:"description"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

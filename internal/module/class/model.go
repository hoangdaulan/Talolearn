package class

import (
	"Talolearn-Project/internal/module/course"
	"Talolearn-Project/internal/module/teacher"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Title       string `gorm:"size:300" json:"title"`
	Description string `json:"description"`

	CourseID  uint `gorm:"not null" json:"course_id"`
	TeacherID uint `gorm:"not null" json:"teacher_id"`

	Course  course.Course   `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Teacher teacher.Teacher `gorm:"foreignKey:TeacherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

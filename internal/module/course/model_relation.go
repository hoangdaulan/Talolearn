package course

import "Talolearn-Project/internal/module/teacher"

type CourseTeacher struct {
	ID        uint `gorm:"primaryKey"`
	CourseID  uint `gorm:"not null;uniqueIndex:idx_course_teacher"`
	TeacherID uint `gorm:"not null;uniqueIndex:idx_course_teacher"`

	Course  Course          `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Teacher teacher.Teacher `gorm:"foreignKey:TeacherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

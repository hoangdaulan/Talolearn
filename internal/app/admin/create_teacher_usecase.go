package admin

import (
	"Talolearn-Project/internal/module/teacher"
	"Talolearn-Project/internal/module/user"

	"gorm.io/gorm"
)

type CreateTeacherUsecase struct {
	teacherRepo teacher.Repository
	userRepo    user.Repository
	db          *gorm.DB
}

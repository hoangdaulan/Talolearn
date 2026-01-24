package teacher

import (
	"github.com/gin-gonic/gin"
)

func TeacherRouter(r *gin.Engine, handler *TeacherHandler) {
	r.POST("/teachers/:id", handler.CreateTeacher)
	r.GET("/teachers/:id", handler.GetTeacherByID)
	r.GET("/teachers/user/:user_id", handler.GetTeacherByUserID)
	r.GET("/teachers/search", handler.FindTeacherByName)
	r.GET("/teachers", handler.GetAllTeachers)
	r.PUT("/teachers/:id", handler.UpdateTeacher)
	r.DELETE("/teachers/:id", handler.DeleteTeacher)
}

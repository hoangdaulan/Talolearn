package course

import "github.com/gin-gonic/gin"

func CourseRouter(r *gin.Engine, handler *CourseHandler) {
	r.POST("/courses", handler.CreateCourse)
	r.GET("/courses/:id", handler.GetCourseByID)
	r.GET("/courses", handler.GetAllCourses)
	r.PUT("/courses/:id", handler.UpdateCourse)
	r.DELETE("/courses/:id", handler.DeleteCourse)
}

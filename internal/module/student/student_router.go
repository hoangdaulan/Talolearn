package student

import "github.com/gin-gonic/gin"

func StudentRouter(r *gin.Engine, h *StudentHandler) {
	studentGroup := r.Group("/students")
	{
		studentGroup.POST("/", h.CreateStudent)
		studentGroup.GET("/:id", h.GetStudent)
		studentGroup.PUT("/:id", h.UpdateStudent)
		studentGroup.DELETE("/:id", h.DeleteStudent)
		studentGroup.GET("/", h.ListStudents)
	}
}

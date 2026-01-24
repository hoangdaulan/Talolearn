package class

import "github.com/gin-gonic/gin"

func ClassRouter(r *gin.Engine, h *ClassHandler) {
	classGroup := r.Group("/classes")
	{
		classGroup.POST("/", h.CreateClass)
		classGroup.GET("/:id", h.GetClass)
		classGroup.PUT("/:id", h.UpdateClass)
		classGroup.DELETE("/:id", h.DeleteClass)
		classGroup.GET("/", h.ListClasses)
	}
}

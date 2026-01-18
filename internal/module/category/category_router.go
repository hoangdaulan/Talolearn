package category

import "github.com/gin-gonic/gin"

func CategoryRouter(r *gin.Engine, handler *CategoryHandler) {
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetCategoryByID)
	r.GET("/category", handler.GetAllCategories)
	r.DELETE("/category/:id", handler.DeleteCategory)
}

package category

import "github.com/gin-gonic/gin"

func CategoryRouter(r *gin.Engine, handler *CategoryHandler) {
	categoryGroup := r.Group("category")
	{
		categoryGroup.POST("", handler.CreateCategory)
		categoryGroup.GET("/:id", handler.GetCategoryByID)
		categoryGroup.PUT("/:id", handler.UpdateCategory)
		categoryGroup.DELETE("/:id", handler.DeleteCategory)
		categoryGroup.GET("", handler.GetAllCategories)
	}
}

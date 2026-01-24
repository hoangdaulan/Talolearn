package user

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine, handler *Handler) {
	users := r.Group("/users")
	{
		users.POST("", handler.CreateUser)
		users.GET("", handler.ListUsers)
		users.GET("/:id", handler.GetUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}

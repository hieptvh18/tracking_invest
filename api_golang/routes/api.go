package routes

import (
	"api_golang/controllers/users"
	"api_golang/repositories"
	"api_golang/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userSvc)

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", userController.List)
			users.GET("/:id", userController.GetByID)
			users.POST("", userController.Create)
			users.PUT("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}

		api.GET("/heathcheck", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"Message" : "Ping pong successful",
			})
		})
	}

	return r
}

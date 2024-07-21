package crud

import (
	"crud-app/internal/app/crud/handler"
	"crud-app/internal/app/crud/repository"
	"crud-app/internal/app/crud/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes initializes all the routes for the application
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/v1")
	{
		api.GET("/health", healthCheck)

		// User routes
		userRoutes := api.Group("/users")
		{
			userRepo := repository.NewGormUserRepository(db)
			userService := service.NewUserService(userRepo)
			userHandler := handler.NewUserHandler(userService)

			userRoutes.GET("/", userHandler.GetAllUsers)
			userRoutes.GET("/:id", userHandler.GetUserByID)
			userRoutes.GET("/username/:username", userHandler.GetUserByUsername)
			userRoutes.POST("/", userHandler.CreateUser)
			userRoutes.PUT("/:id", userHandler.UpdateUser)
			userRoutes.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

package main

import (
	"clean-code/controller"
	"clean-code/database"
	router "clean-code/http"
	"clean-code/repository"
	"clean-code/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func runApp(gin *gin.Engine, gorm *gorm.DB) {
	// migrate db
	database.Migrate(gorm)

	r := router.NewBaseRouter(&gin.RouterGroup)

	// users
	userRepo := repository.NewUserRepository(gorm)
	userService := service.NewUserService(userRepo)
	controller.NewUserController(userService, r)

}

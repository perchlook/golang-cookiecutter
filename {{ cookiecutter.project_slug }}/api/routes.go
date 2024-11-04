package api

import (
	"github.com/candorship/candorship/repositories"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, userRepo repositories.UserRepository, config *oauth2.Config) {
	handlers := NewUserHandlers(userRepo, config)

	e.GET("/users", handlers.CreateUser)
}

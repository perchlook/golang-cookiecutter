package api

import (
	"github.com/candorship/candorship/repositories"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type UserHandlers struct {
	userRepository repositories.UserRepository
	oauthConfig    *oauth2.Config
}

func NewUserHandlers(userRepository repositories.UserRepository, oauthConfig *oauth2.Config) *UserHandlers {
	return &UserHandlers{userRepository, oauthConfig}
}

func (h *UserHandlers) GetUser(c echo.Context) error {
	username := c.Param("username")
	user, err := h.userRepository.GetUser(username)

	if err != nil {
		return c.String(404, "User cannot be found")
	}
	return c.JSON(200, user)
}

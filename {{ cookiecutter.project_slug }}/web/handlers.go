package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/models"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/repositories"
	"golang.org/x/oauth2"
	oauthv2 "google.golang.org/api/oauth2/v2"
	option "google.golang.org/api/option"
)

type WebHandlers struct {
	userRepository repositories.UserRepository
	oauthConfig    *oauth2.Config
}

func NewWebHandlers(userRepository repositories.UserRepository, oauthConfig *oauth2.Config) *WebHandlers {
	return &WebHandlers{userRepository, oauthConfig}
}

func (h *WebHandlers) IndexPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (h *WebHandlers) GoogleLogin(c echo.Context) error {
	url := h.oauthConfig.AuthCodeURL("")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *WebHandlers) GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	ctx := c.Request().Context()

	token, err := h.oauthConfig.Exchange(ctx, code)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	service, err := oauthv2.NewService(ctx, option.WithTokenSource(h.oauthConfig.TokenSource(ctx, token)))

	if err != nil {
		err := fmt.Errorf("error getting oauth2 service %v", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resp, err := service.Userinfo.Get().Do()
	if err != nil {
		err := fmt.Errorf("error getting user info %v", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	email := resp.Email
	user, err := h.userRepository.FindByEmail(email)

	if err != nil {
		fmt.Println("Create user with google profile")
		user = &models.User{Email: email, Username: email}
		profile := &models.GoogleProfile{Email: email, Name: resp.Name, OAuthToken: token.AccessToken, RefreshToken: token.RefreshToken, TokenExpiry: token.Expiry, Scopes: nil}
		user, err = h.userRepository.CreateUserWithGoogleProfile(user, profile)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

	} else {
		fmt.Println("Existing user")
		fmt.Println("Saved google profile ID: ", user.GoogleProfile.ID)
	}

	fmt.Println(user.ID)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

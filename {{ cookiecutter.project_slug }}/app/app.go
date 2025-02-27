package app

import (
	"html/template"
	"io"
	"os"

	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/labstack/echo/v4"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/api"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/database"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/models"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/repositories"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/web"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InitApp(databaseUrl string) *echo.Echo {
	db := database.InitDB(databaseUrl)
	db.AutoMigrate(&models.User{}, &models.GoogleProfile{})
	userRepo := repositories.NewUserRepository(db)

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Renderer = t
	e.Static("/static", "static")

	api.RegisterRoutes(e, db, userRepo, conf)
	web.RegisterRoutes(e, db, userRepo, conf)

	return e
}

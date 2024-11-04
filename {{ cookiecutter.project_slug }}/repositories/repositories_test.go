package repositories

import (
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/database"
	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/models"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Use temporary file for test database
	tmpDB := ":memory:"
	db := database.InitDB(tmpDB)
	db.AutoMigrate(&models.User{}, &models.GoogleProfile{})
	return db
}

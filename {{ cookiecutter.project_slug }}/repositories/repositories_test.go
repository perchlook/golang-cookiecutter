package repositories

import (
	"github.com/candorship/candorship/database"
	"github.com/candorship/candorship/models"
	"gorm.io/gorm"
)


func setupTestDB() *gorm.DB {
	// Use temporary file for test database
	tmpDB := ":memory:"
	db := database.InitDB(tmpDB)
	db.AutoMigrate(&models.User{}, &models.Organization{}, &models.Goal{}, &models.GoogleProfile{})
	return db
}
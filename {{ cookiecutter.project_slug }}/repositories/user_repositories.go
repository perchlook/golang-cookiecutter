package repositories

import (
	"errors"

	"gorm.io/gorm"

	"github.com/{{ cookiecutter.org_name }}/{{ cookiecutter.project_slug }}/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	CreateUserWithGoogleProfile(user *models.User, googleProfile *models.GoogleProfile) (*models.User, error)
	GetUser(username string) (*models.User, error)
	CheckUsernameExists(username string) bool
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if r.CheckUsernameExists(user.Username) {
		return nil, errors.New("username has been taken")
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *userRepository) CreateUserWithGoogleProfile(user *models.User, googleProfile *models.GoogleProfile) (*models.User, error) {
	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if r.CheckUsernameExists(user.Username) {
		return nil, errors.New("username has been taken")
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}

	googleProfile.UserID = user.ID
	result = r.db.Create(&googleProfile)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *userRepository) GetUser(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	user := models.User{}
	r.db.Select("id, username, created_at").Where("username = ?", username).First(&user)
	return &user, nil
}

func (r *userRepository) CheckUsernameExists(username string) bool {
	var count int64

	r.db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	user := models.User{}
	result := r.db.Preload("GoogleProfile").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

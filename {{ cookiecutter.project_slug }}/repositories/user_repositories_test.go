package repositories

import (
	"testing"
	"time"

	"github.com/candorship/candorship/models"
)

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	userRepo := NewUserRepository(db)

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		{
			name:    "Successfully create user",
			user:    &models.User{Username: "testuser", Email: "a@b.com"},
			wantErr: false,
		},
		{
			name:    "Successfully create user with organization",
			user:    &models.User{Username: "testuser2", Email: "a@c.com"},
			wantErr: false,
		},
		{
			name:    "Create user with empty username",
			user:    &models.User{Username: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			user, err := userRepo.CreateUser(tt.user)

			// Verify
			if tt.wantErr {
				if err == nil {
					t.Error("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if user.ID == "" {
					t.Error("expected non-empty user ID")
				}
				if user.Username != tt.user.Username {
					t.Errorf("expected username %v, got %v", tt.user.Username, user.Username)
				}
				if time.Since(user.CreatedAt) > time.Second {
					t.Error("CreatedAt time is too old")
				}
			}
		})
	}
}

func TestCheckUsernameExists(t *testing.T) {
	db := setupTestDB()
	userRepo := NewUserRepository(db)

	// Create a user to test against
	user := &models.User{Username: "existinguser", Email: "existing@a.com"}
	_, err := userRepo.CreateUser(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	tests := []struct {
		name     string
		username string
		exists   bool
	}{
		{
			name:     "Username exists",
			username: "existinguser",
			exists:   true,
		},
		{
			name:     "Username does not exist",
			username: "nonexistentuser",
			exists:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			exists := userRepo.CheckUsernameExists(tt.username)

			// Verify
			if exists != tt.exists {
				t.Errorf("expected %v, got %v", tt.exists, exists)
			}
		})
	}
}

func TestCreateUserWithGoogleProfile(t *testing.T) {
	db := setupTestDB()
	userRepo := NewUserRepository(db)

	user := &models.User{Username: "testuser", Email: "google@profile.com"}

	tests := []struct {
		name          string
		user          *models.User
		googleProfile *models.GoogleProfile
		wantErr       bool
	}{
		{
			name:          "Successfully create user with google profile",
			user:          user,
			googleProfile: &models.GoogleProfile{Email: "google@profile.com", Name: "Test", OAuthToken: "token", RefreshToken: "refresh", TokenExpiry: time.Now(), Scopes: []string{}},
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			user, err := userRepo.CreateUserWithGoogleProfile(tt.user, tt.googleProfile)

			// Verify
			if tt.wantErr {
				if err == nil {
					t.Error("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if user.ID == "" {
					t.Error("expected non-empty user ID")
				}
				profile := &models.GoogleProfile{}
				db.Where("user_id = ?", user.ID).First(profile)
				if profile.UserID != user.ID {
					t.Errorf("expected user ID %v, got %v", user.ID, profile.UserID)
				}
			}

		})
	}
}

func TestFindByEmail(t *testing.T) {
	db := setupTestDB()
	userRepo := NewUserRepository(db)

	// Create a user to test against
	user := &models.User{Username: "testuser", Email: "test@example.com"}
	_, err := userRepo.CreateUser(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "Email exists",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "Email does not exist",
			email:   "nonexistent@example.com",
			wantErr: true,
		},
		{
			name:    "Empty email",
			email:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := userRepo.FindByEmail(tt.email)

			if tt.wantErr {
				if err == nil {
					t.Error("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if user.Email != tt.email {
					t.Errorf("expected email %v, got %v", tt.email, user.Email)
				}
			}
		})
	}
}

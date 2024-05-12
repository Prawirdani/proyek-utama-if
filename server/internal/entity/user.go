package entity

import (
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
	"github.com/prawirdani/golang-restapi/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorWrongCredentials = httputil.ErrUnauthorized("Check your credentials")
)

type UserRole string

const (
	Kasir   UserRole = "Kasir"
	Manajer UserRole = "Manajer"
)

type User struct {
	ID        int        `json:"id,omitempty"`
	Nama      string     `json:"nama" validate:"required"`
	Username  string     `json:"username,omitempty" validate:"required"`
	Password  string     `json:"-" validate:"required,min=6"`
	Active    bool       `json:"active,omitempty"`
	Role      UserRole   `json:"role,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Create new user from request payload
func NewUser(request model.RegisterRequest) User {
	return User{
		Nama:     request.Nama,
		Username: request.Username,
		Password: request.Password,
	}
}

// Semantic Validation
func (u User) Validate() error {
	return utils.Validate.Struct(u)
}

// Encrypt user password
func (u *User) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Verify / Decrypt user password
func (u User) VerifyPassword(plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	if err != nil {
		return ErrorWrongCredentials
	}
	return nil
}

// Generate JWT Token
func (u User) GenerateToken(secret string, expiryHour int) (string, error) {
	payload := utils.NewJwtClaims(
		map[string]interface{}{
			"id":       u.ID,
			"nama":     u.Nama,
			"username": u.Username,
			"role":     u.Role,
		},
		"user",
	)

	return utils.GenerateToken(payload, secret, time.Duration(expiryHour)*time.Hour)
}

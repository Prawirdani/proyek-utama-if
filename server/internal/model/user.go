package model

import "time"

type UserUpdateRequest struct {
	ID       int
	Nama     string `json:"nama" validate:"required"`
	Username string `json:"username,omitempty" validate:"required"`
}

type UserResetPasswordRequest struct {
	ID          int
	NewPassword string `json:"newPassword" validate:"required"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Username  string    `json:"username"`
	Active    bool      `json:"active"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

package model

import "github.com/prawirdani/golang-restapi/pkg/utils"

type RegisterRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Syntatic Validator
func (r RegisterRequest) ValidateRequest() error {
	return utils.Validate.Struct(r)
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Syntatic Validator
func (r LoginRequest) ValidateRequest() error {
	return utils.Validate.Struct(r)
}

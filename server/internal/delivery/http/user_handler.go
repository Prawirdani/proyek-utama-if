package http

import (
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/usecase"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

type UserHandler struct {
	cfg    *config.Config
	userUC usecase.UserUseCase
}

func NewUserHandler(cfg *config.Config, userUC usecase.UserUseCase) UserHandler {
	return UserHandler{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (h UserHandler) HandleListUser(w http.ResponseWriter, r *http.Request) error {
	users, err := h.userUC.ListUser(r.Context())
	if err != nil {
		return err
	}

	resData := ToResponseList(users, entity.User.ToResponse)
	return response(w, status(http.StatusOK), data(resData))
}

func (h UserHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "userID")
	if err != nil {
		return err
	}

	reqBody, err := BindAndValidate[model.UserUpdateRequest](r)
	if err != nil {
		return err
	}
	reqBody.ID = id

	if err := h.userUC.UpdateUser(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Sukses update user."))
}

func (h UserHandler) HandleActivateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "userID")
	if err != nil {
		return err
	}

	if err := h.userUC.ActivateUser(r.Context(), id); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Sukses mengaktifkan user."))
}

func (h UserHandler) HandleDeactivateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "userID")
	if err != nil {
		return err
	}

	if err := h.userUC.DeactivateUser(r.Context(), id); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Sukses menonaktifkan user."))
}

func (h UserHandler) HandleResetPassword(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "userID")
	if err != nil {
		return err
	}
	reqBody, err := BindAndValidate[model.UserResetPasswordRequest](r)
	if err != nil {
		return err
	}
	reqBody.ID = id

	if err := h.userUC.ResetPassword(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Sukses mereset password user."))
}

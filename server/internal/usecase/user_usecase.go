package usecase

import (
	"context"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type UserUseCase interface {
	ListUser(ctx context.Context) ([]entity.User, error)
	DeactivateUser(ctx context.Context, id int) error
	ActivateUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, request model.UserUpdateRequest) error
	ResetPassword(ctx context.Context, request model.UserResetPasswordRequest) error
}
type userUseCase struct {
	cfg      *config.Config
	userRepo repository.UserRepository
}

func NewUserUseCase(cfg *config.Config, userRepo repository.UserRepository) userUseCase {
	return userUseCase{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (u userUseCase) ListUser(ctx context.Context) ([]entity.User, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()
	return u.userRepo.Select(ctxWT)
}

func (u userUseCase) UpdateUser(ctx context.Context, request model.UserUpdateRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	user, err := u.userRepo.SelectWhere(ctxWT, "id", request.ID)
	if err != nil {
		return err
	}
	user.AssignUpdate(request)
	return u.userRepo.Update(ctxWT, user)
}

func (u userUseCase) DeactivateUser(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	user, err := u.userRepo.SelectWhere(ctxWT, "id", id)
	if err != nil {
		return err
	}
	user.Deactivate()
	return u.userRepo.Update(ctxWT, user)
}

func (u userUseCase) ActivateUser(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	user, err := u.userRepo.SelectWhere(ctxWT, "id", id)
	if err != nil {
		return err
	}
	user.Activate()
	return u.userRepo.Update(ctxWT, user)
}

func (u userUseCase) ResetPassword(ctx context.Context, request model.UserResetPasswordRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	user, err := u.userRepo.SelectWhere(ctxWT, "id", request.ID)
	if err != nil {
		return err
	}
	user.NewPassword(request.NewPassword)
	if err := user.EncryptPassword(); err != nil {
		return err
	}

	return u.userRepo.Update(ctxWT, user)
}

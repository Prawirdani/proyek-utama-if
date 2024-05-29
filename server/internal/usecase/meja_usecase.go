package usecase

import (
	"context"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type MejaUseCase interface {
	CreateMeja(ctx context.Context, request model.CreateMejaRequest) error
	ListMeja(ctx context.Context) ([]entity.Meja, error)
	FindMeja(ctx context.Context, id int) (*entity.Meja, error)
	UpdateMeja(ctx context.Context, request model.UpdateMejaRequest) error
	RemoveMeja(ctx context.Context, id int) error
}

type mejaUseCase struct {
	mejaRepo repository.MejaRepository
	cfg      *config.Config
}

func NewMejaUseCase(cfg *config.Config, mejaRepo repository.MejaRepository) mejaUseCase {
	return mejaUseCase{
		mejaRepo: mejaRepo,
		cfg:      cfg,
	}
}

func (us mejaUseCase) CreateMeja(ctx context.Context, request model.CreateMejaRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	newMeja := entity.NewMeja(request)

	return us.mejaRepo.Insert(ctxWT, *newMeja)
}

func (us mejaUseCase) ListMeja(ctx context.Context) ([]entity.Meja, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	return us.mejaRepo.Select(ctxWT)
}

func (us mejaUseCase) FindMeja(ctx context.Context, id int) (*entity.Meja, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	return us.mejaRepo.SelectWhere(ctxWT, "id", id)
}

func (us mejaUseCase) UpdateMeja(ctx context.Context, request model.UpdateMejaRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	meja, err := us.mejaRepo.SelectWhere(ctxWT, "id", request.ID)
	if err != nil {
		return err
	}
	meja.Assign(request)

	return us.mejaRepo.Update(ctxWT, *meja)
}

func (us mejaUseCase) RemoveMeja(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	meja, err := us.mejaRepo.SelectWhere(ctxWT, "id", id)
	if err != nil {
		return err
	}
	meja.SetDeletedAt()

	return us.mejaRepo.Update(ctxWT, *meja)
}

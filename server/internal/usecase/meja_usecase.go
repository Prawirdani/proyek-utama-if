package usecase

import (
	"context"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type MejaUseCase interface {
	CreateMeja(ctx context.Context, request model.CreateMejaRequest) error
	ListMeja(ctx context.Context) ([]entity.Meja, error)
	FindMeja(ctx context.Context, id int) (entity.Meja, error)
	UpdateMeja(ctx context.Context, request model.UpdateMejaRequest) error
	RemoveMeja(ctx context.Context, id int) error
}

type mejaUseCase struct {
	mejaRepo repository.MejaRepository
	cfg      *config.Config
}

func NewMejaUseCase(mejaRepo repository.MejaRepository, cfg *config.Config) mejaUseCase {
	return mejaUseCase{
		mejaRepo: mejaRepo,
		cfg:      cfg,
	}
}

func (us mejaUseCase) CreateMeja(ctx context.Context, request model.CreateMejaRequest) error {
	newMeja := entity.NewMeja(request)

	return us.mejaRepo.Insert(ctx, newMeja)
}

func (us mejaUseCase) ListMeja(ctx context.Context) ([]entity.Meja, error) {
	return us.mejaRepo.Select(ctx)
}

func (us mejaUseCase) FindMeja(ctx context.Context, id int) (entity.Meja, error) {
	return us.mejaRepo.SelectWhere(ctx, "id", id)
}

func (us mejaUseCase) UpdateMeja(ctx context.Context, request model.UpdateMejaRequest) error {
	return nil
}

func (us mejaUseCase) RemoveMeja(ctx context.Context, id int) error {
	return nil
}

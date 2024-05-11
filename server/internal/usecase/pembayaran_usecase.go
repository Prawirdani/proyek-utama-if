package usecase

import (
	"context"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type PembayaranUseCase interface {
	CreateMetodePembayaran(ctx context.Context, request model.CreateMetodePembayaranRequest) error
	ListMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error)
	FindMetodePembayaran(ctx context.Context, id int) (*entity.MetodePembayaran, error)
	UpdateMetodePembayaran(ctx context.Context, request model.UpdateMetodePembayaranRequest) error
	RemoveMetodePembayaran(ctx context.Context, id int) error
}

type pembayaranUsecase struct {
	pembayaranRepo repository.PembayaranRepository
	cfg            *config.Config
}

func NewPembayaranUsecase(pembayaranRepo repository.PembayaranRepository, cfg *config.Config) pembayaranUsecase {
	return pembayaranUsecase{
		pembayaranRepo: pembayaranRepo,
		cfg:            cfg,
	}
}

func (u pembayaranUsecase) CreateMetodePembayaran(ctx context.Context, request model.CreateMetodePembayaranRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := entity.NewMetodePembayaran(request)
	if err != nil {
		return err
	}

	err = u.pembayaranRepo.InsertMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

func (u pembayaranUsecase) ListMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mps, err := u.pembayaranRepo.SelectMetodePembayaran(ctxWT)
	if err != nil {
		return nil, err
	}

	return mps, nil
}

func (u pembayaranUsecase) FindMetodePembayaran(ctx context.Context, id int) (*entity.MetodePembayaran, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.pembayaranRepo.SelectMetodePembayaranWhere(ctxWT, "id", id)
	if err != nil {
		return nil, err
	}

	return mp, nil
}

func (u pembayaranUsecase) UpdateMetodePembayaran(ctx context.Context, request model.UpdateMetodePembayaranRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.FindMetodePembayaran(ctxWT, request.ID)
	if err != nil {
		return err
	}

	err = mp.Assign(request)
	if err != nil {
		return err
	}

	err = u.pembayaranRepo.UpdateMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

func (u pembayaranUsecase) RemoveMetodePembayaran(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.FindMetodePembayaran(ctxWT, id)
	if err != nil {
		return err
	}

	mp.SetDeletedAt()
	err = u.pembayaranRepo.UpdateMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

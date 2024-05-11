package usecase

import (
	"context"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
	"github.com/prawirdani/golang-restapi/internal/valueobject"
)

type PaymentUsecase interface {
	CreateMetodePembayaran(ctx context.Context, request model.CreateMetodePembayaranRequest) error
	ListMetodePembayaran(ctx context.Context) ([]valueobject.MetodePembayaran, error)
	FindMetodePembayaran(ctx context.Context, id int) (*valueobject.MetodePembayaran, error)
	UpdateMetodePembayaran(ctx context.Context, request model.UpdateMetodePembayaranRequest) error
	RemoveMetodePembayaran(ctx context.Context, id int) error
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
	cfg         *config.Config
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository, cfg *config.Config) paymentUsecase {
	return paymentUsecase{
		paymentRepo: paymentRepo,
		cfg:         cfg,
	}
}

func (u paymentUsecase) CreateMetodePembayaran(ctx context.Context, request model.CreateMetodePembayaranRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := valueobject.NewMetodePembayaran(request)
	if err != nil {
		return err
	}

	err = u.paymentRepo.InsertMetodePembayaran(ctxWT, mp)
	if err != nil {
		return err
	}

	return nil
}

func (u paymentUsecase) ListMetodePembayaran(ctx context.Context) ([]valueobject.MetodePembayaran, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mps, err := u.paymentRepo.SelectMetodePembayaran(ctxWT)
	if err != nil {
		return nil, err
	}

	return mps, nil
}

func (u paymentUsecase) FindMetodePembayaran(ctx context.Context, id int) (*valueobject.MetodePembayaran, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.paymentRepo.SelectMetodePembayaranWhere(ctxWT, "id", id)
	if err != nil {
		return nil, err
	}

	return mp, nil
}

func (u paymentUsecase) UpdateMetodePembayaran(ctx context.Context, request model.UpdateMetodePembayaranRequest) error {
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

	err = u.paymentRepo.UpdateMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

func (u paymentUsecase) RemoveMetodePembayaran(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.FindMetodePembayaran(ctxWT, id)
	if err != nil {
		return err
	}

	mp.SetDeletedAt()
	err = u.paymentRepo.UpdateMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

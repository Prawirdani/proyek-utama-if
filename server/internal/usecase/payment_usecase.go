package usecase

import (
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type PaymentUsecase interface {
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
	cfg         *config.Config
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository, cfg *config.Config) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		cfg:         cfg,
	}
}

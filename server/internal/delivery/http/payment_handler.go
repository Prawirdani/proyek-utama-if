package http

import (
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/usecase"
)

type PaymentHandler struct {
	paymentUC usecase.PaymentUsecase
	cfg       *config.Config
}

func NewPaymentHandler(paymentUC usecase.PaymentUsecase, cfg *config.Config) PaymentHandler {
	return PaymentHandler{
		paymentUC: paymentUC,
		cfg:       cfg,
	}
}

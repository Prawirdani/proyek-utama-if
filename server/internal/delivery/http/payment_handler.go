package http

import (
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/usecase"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
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

func (h PaymentHandler) HandleCreateMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := BindAndValidate[model.CreateMetodePembayaranRequest](r)
	if err != nil {
		return err
	}

	if err := h.paymentUC.CreateMetodePembayaran(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Metode pembayaran berhasil ditambahkan!"))
}

func (h PaymentHandler) HandleListMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	mps, err := h.paymentUC.ListMetodePembayaran(r.Context())
	if err != nil {
		return err
	}

	return response(w, status(http.StatusOK), data(mps))
}

func (h PaymentHandler) HandleFindMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "metodePembayaranID")

	mp, err := h.paymentUC.FindMetodePembayaran(r.Context(), id)
	if err != nil {
		return err
	}

	return response(w, status(http.StatusOK), data(mp))
}

func (h PaymentHandler) HandleUpdateMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "metodePembayaranID")
	if err != nil {
		return err
	}

	reqBody, err := BindAndValidate[model.UpdateMetodePembayaranRequest](r)
	if err != nil {
		return err
	}
	reqBody.ID = id

	if err := h.paymentUC.UpdateMetodePembayaran(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Metode pembayaran berhasil diupdate!"))
}

func (h PaymentHandler) HandleDeleteMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "metodePembayaranID")
	if err != nil {
		return err
	}

	if err := h.paymentUC.RemoveMetodePembayaran(r.Context(), id); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Metode pembayaran berhasil dihapus!"))
}

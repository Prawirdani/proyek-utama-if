package http

import (
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/usecase"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

type PembayaranHandler struct {
	pembayaranUC usecase.PembayaranUseCase
	cfg          *config.Config
}

func NewPembayaranHandler(cfg *config.Config, pembayaranUC usecase.PembayaranUseCase) PembayaranHandler {
	return PembayaranHandler{
		pembayaranUC: pembayaranUC,
		cfg:          cfg,
	}
}

func (h PembayaranHandler) HandleCreateMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := BindAndValidate[model.CreateMetodePembayaranRequest](r)
	if err != nil {
		return err
	}

	if err := h.pembayaranUC.CreateMetodePembayaran(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Metode pembayaran berhasil ditambahkan!"))
}

func (h PembayaranHandler) HandleListMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	mps, err := h.pembayaranUC.ListMetodePembayaran(r.Context())
	if err != nil {
		return err
	}

	return response(w, status(http.StatusOK), data(mps))
}

func (h PembayaranHandler) HandleFindMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "metodePembayaranID")

	mp, err := h.pembayaranUC.FindMetodePembayaran(r.Context(), id)
	if err != nil {
		return err
	}

	return response(w, status(http.StatusOK), data(mp))
}

func (h PembayaranHandler) HandleUpdateMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "metodePembayaranID")
	if err != nil {
		return err
	}

	reqBody, err := BindAndValidate[model.UpdateMetodePembayaranRequest](r)
	if err != nil {
		return err
	}
	reqBody.ID = id

	if err := h.pembayaranUC.UpdateMetodePembayaran(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Metode pembayaran berhasil diupdate!"))
}

func (h PembayaranHandler) HandleDeleteMetodePembayaran(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "metodePembayaranID")
	if err != nil {
		return err
	}

	if err := h.pembayaranUC.RemoveMetodePembayaran(r.Context(), id); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Metode pembayaran berhasil dihapus!"))
}

func (h PembayaranHandler) HandleBayarPesanan(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := BindAndValidate[model.PembayaranRequest](r)
	if err != nil {
		return err
	}

	if err := h.pembayaranUC.BayarPesanan(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Pembayaran berhasil!"))
}

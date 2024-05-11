package http

import (
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/usecase"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

type MejaHandler struct {
	mejaUC usecase.MejaUseCase
	cfg    *config.Config
}

func NewMejaHandler(cfg *config.Config, mejaUC usecase.MejaUseCase) MejaHandler {
	return MejaHandler{
		mejaUC: mejaUC,
		cfg:    cfg,
	}
}

func (h MejaHandler) HandleCreateMeja(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := BindAndValidate[model.CreateMejaRequest](r)
	if err != nil {
		return err
	}

	if err := h.mejaUC.CreateMeja(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Meja berhasil dibuat!"))
}

func (h MejaHandler) HandleListMeja(w http.ResponseWriter, r *http.Request) error {
	mejas, err := h.mejaUC.ListMeja(r.Context())
	if err != nil {
		return err
	}
	return response(w, data(mejas))
}

func (h MejaHandler) HandleFindMeja(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "tableID")
	if err != nil {
		return httputil.ErrBadRequest("ID meja tidak valid!")
	}

	meja, err := h.mejaUC.FindMeja(r.Context(), id)
	if err != nil {
		return err
	}

	return response(w, data(meja))
}

func (h MejaHandler) HandleUpdateMeja(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "tableID")
	if err != nil {
		return httputil.ErrBadRequest("ID meja tidak valid!")
	}

	reqBody, err := BindAndValidate[model.UpdateMejaRequest](r)
	if err != nil {
		return err
	}
	reqBody.ID = id

	if err := h.mejaUC.UpdateMeja(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, message("Meja berhasil diupdate!"))
}

func (h MejaHandler) HandleDeleteMeja(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "tableID")
	if err != nil {
		return httputil.ErrBadRequest("ID meja tidak valid!")
	}

	if err := h.mejaUC.RemoveMeja(r.Context(), id); err != nil {
		return err
	}

	return response(w, message("Meja berhasil dihapus!"))
}

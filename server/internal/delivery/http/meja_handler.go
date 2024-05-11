package http

import (
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/usecase"
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
	return nil
}

func (h MejaHandler) HandleListMeja(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h MejaHandler) HandleFindMeja(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h MejaHandler) HandleUpdateMeja(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h MejaHandler) HandleDeleteMeja(w http.ResponseWriter, r *http.Request) error {
	return nil
}

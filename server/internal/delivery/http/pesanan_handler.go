package http

import (
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/usecase"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

type PesananHandler struct {
	cfg       *config.Config
	pesananUC usecase.PesananUseCase
}

func NewPesananHandler(cfg *config.Config, pesananUC usecase.PesananUseCase) PesananHandler {
	return PesananHandler{
		cfg:       cfg,
		pesananUC: pesananUC,
	}
}

func (ph PesananHandler) HandlePesananDineIn(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := BindAndValidate[model.PesananDineInRequest](r)
	if err != nil {
		return err
	}

	authClaims := httputil.GetAuthCtx(r.Context())
	userId := int(authClaims["user"].(map[string]interface{})["id"].(float64))
	reqBody.KasirID = userId

	if err := ph.pesananUC.CreateDineIn(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Pesanan dine-in berhasil dibuat"))
}

func (ph PesananHandler) HandlePesananTakeAway(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := BindAndValidate[model.PesananTakeAwayRequest](r)
	if err != nil {
		return err
	}

	authClaims := httputil.GetAuthCtx(r.Context())
	userId := int(authClaims["user"].(map[string]interface{})["id"].(float64))
	reqBody.KasirID = userId

	if err := ph.pesananUC.CreateTakeAway(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Pesanan take-away berhasil dibuat"))
}

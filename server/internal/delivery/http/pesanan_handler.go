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

func (ph PesananHandler) HandleListPesanan(w http.ResponseWriter, r *http.Request) error {
	ps, err := ph.pesananUC.ListPesanan(r.Context())
	if err != nil {
		return err
	}

	return response(w, status(http.StatusOK), data(ps))
}

func (ph PesananHandler) HandleFindPesanan(w http.ResponseWriter, r *http.Request) error {
	pesananID, err := httputil.ParamInt(r, "pesananID")
	if err != nil {
		return err
	}

	pesanan, err := ph.pesananUC.FindPesanan(r.Context(), pesananID)
	if err != nil {
		return err
	}

	return response(w, status(http.StatusOK), data(pesanan))
}

func (ph PesananHandler) HandleAddMenu(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := BindAndValidate[model.PesananMenuRequest](r)
	if err != nil {
		return err
	}

	pesananID, err := httputil.ParamInt(r, "pesananID")
	if err != nil {
		return err
	}

	if err := ph.pesananUC.AddMenuToPesanan(r.Context(), pesananID, reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Menu berhasil ditambahkan ke pesanan"))
}

func (ph PesananHandler) HandleRemoveMenu(w http.ResponseWriter, r *http.Request) error {
	pesananID, err := httputil.ParamInt(r, "pesananID")
	if err != nil {
		return err
	}

	detailID, err := httputil.ParamInt(r, "detailID")
	if err != nil {
		return err
	}

	if err := ph.pesananUC.RemoveMenuFromPesanan(r.Context(), pesananID, detailID); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Menu berhasil dihapus dari pesanan"))
}

func (ph PesananHandler) HandleBatalkanPesanan(w http.ResponseWriter, r *http.Request) error {
	pesananID, err := httputil.ParamInt(r, "pesananID")
	if err != nil {
		return err
	}

	if err := ph.pesananUC.BatalkanPesanan(r.Context(), pesananID); err != nil {
		return err
	}

	return response(w, status(http.StatusOK), message("Pesanan berhasil dibatalkan"))
}
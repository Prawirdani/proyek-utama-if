package http

import (
	"encoding/json"
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/usecase"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
	"github.com/prawirdani/golang-restapi/pkg/utils"
)

type MenuHandler struct {
	menuUC usecase.MenuUsecase
	cfg    *config.Config
}

func NewMenuHandler(cfg *config.Config, menuUC usecase.MenuUsecase) MenuHandler {
	return MenuHandler{
		menuUC: menuUC,
		cfg:    cfg,
	}
}

func (h MenuHandler) HandleCreateKategori(w http.ResponseWriter, r *http.Request) error {
	var reqBody model.CreateKategoriMenuRequest

	if err := httputil.BindJSON(r, &reqBody); err != nil {
		return err
	}

	if err := utils.Validate.Struct(reqBody); err != nil {
		return err
	}

	if err := h.menuUC.CreateKategori(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(201), message("Kategori menu created."))
}

func (h MenuHandler) HandleCreateMenu(w http.ResponseWriter, r *http.Request) error {
	dataBody := r.FormValue("data")
	var menuData model.CreateMenuRequest

	err := json.Unmarshal([]byte(dataBody), &menuData)
	if err != nil {
		return httputil.ErrBadRequest(err.Error())
	}

	if err := utils.Validate.Struct(menuData); err != nil {
		return err
	}

	imageName, err := httputil.UploadHandler(r, "image")
	if err != nil {
		return err
	}

	menuData.ImageName = imageName
	if err := h.menuUC.CreateMenu(r.Context(), menuData); err != nil {
		httputil.DeleteUpload(*imageName)
		return err
	}
	return response(w, status(201), message("Menu created."))
}

func (h MenuHandler) HandleListKategori(w http.ResponseWriter, r *http.Request) error {
	kategori, err := h.menuUC.ListKategori(r.Context())
	if err != nil {
		return err
	}

	return response(w, status(200), data(kategori))
}

func (h MenuHandler) HandleListMenu(w http.ResponseWriter, r *http.Request) error {
	menu, err := h.menuUC.ListMenu(r.Context())
	if err != nil {
		return err
	}

	return response(w, status(200), data(menu))
}

func (h MenuHandler) HandleFindMenu(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "menuID")
	if err != nil {
		return httputil.ErrBadRequest("Invalid menu ID.")
	}

	menu, err := h.menuUC.FindMenu(r.Context(), id)
	if err != nil {
		return err
	}
	menu.FormatURL(h.cfg)

	return response(w, status(200), data(menu))
}

func (h MenuHandler) HandleDeleteKategori(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "categoryID")
	if err != nil {
		return httputil.ErrBadRequest("Invalid category ID.")
	}

	if err := h.menuUC.RemoveKategori(r.Context(), id); err != nil {
		return err
	}

	return response(w, status(200), message("Kategori menu removed."))
}

func (h MenuHandler) HandleDeleteMenu(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "menuID")
	if err != nil {
		return httputil.ErrBadRequest("Invalid menu ID.")
	}

	if err := h.menuUC.RemoveMenu(r.Context(), id); err != nil {
		return err
	}

	return response(w, status(200), message("Menu removed."))
}

func (h MenuHandler) HandleUpdateKategori(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "categoryID")
	if err != nil {
		return httputil.ErrBadRequest("Invalid category ID.")
	}

	var reqBody model.UpdateKategoriMenuRequest
	if err := httputil.BindJSON(r, &reqBody); err != nil {
		return err
	}
	reqBody.ID = id

	if err := utils.Validate.Struct(reqBody); err != nil {
		return err
	}

	if err := h.menuUC.UpdateKategori(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(200), message("Kategori menu updated."))
}

func (h MenuHandler) HandleUpdateMenu(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "menuID")
	if err != nil {
		return httputil.ErrBadRequest("Invalid menu ID.")
	}

	dataBody := r.FormValue("data")
	var menuData model.UpdateMenuRequest

	err = json.Unmarshal([]byte(dataBody), &menuData)
	if err != nil {
		return httputil.ErrBadRequest(err.Error())
	}

	menuData.ID = id
	if err := utils.Validate.Struct(menuData); err != nil {
		return err
	}

	imageName, err := httputil.UploadHandler(r, "image")
	if err != nil {
		return err
	}
	menuData.ImageName = imageName

	if err := h.menuUC.UpdateMenu(r.Context(), menuData); err != nil {
		httputil.DeleteUpload(*imageName)
		return err
	}

	return response(w, status(200), message("Menu updated."))
}

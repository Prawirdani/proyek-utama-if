package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
	reqBody, err := BindAndValidate[model.CreateKategoriMenuRequest](r)
	if err != nil {
		return err
	}

	if err := h.menuUC.CreateKategori(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(http.StatusCreated), message("Kategori menu berhasil dibuat!."))
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

	file, handler, err := r.FormFile("image")

	if err != nil {
		slog.Error("UploadHandler.FormFile", slog.String("details", err.Error()))
		return err
	}
	defer file.Close()

	extension := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	// Create a new file in the server's upload directory
	// /project/uploads
	f, err := os.OpenFile(filepath.Join("uploads", newFileName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		slog.Error("UploadHandler.OpenFile", slog.String("details", err.Error()))
		return err
	}
	defer f.Close()

	// Copy the file to the destination path
	_, err = io.Copy(f, file)
	if err != nil {
		slog.Error("UploadHandler.Copy", slog.String("details:", err.Error()))
		return err
	}

	menuData.ImageName = &newFileName
	if err := h.menuUC.CreateMenu(r.Context(), menuData); err != nil {
		httputil.DeleteUpload(*menuData.ImageName)
		return err
	}

	return response(w, status(http.StatusCreated), message("Menu berhasil dibuat!."))
}

func (h MenuHandler) HandleListKategori(w http.ResponseWriter, r *http.Request) error {
	kategori, err := h.menuUC.ListKategori(r.Context())
	if err != nil {
		return err
	}

	return response(w, data(kategori))
}

func (h MenuHandler) HandleListMenu(w http.ResponseWriter, r *http.Request) error {
	menu, err := h.menuUC.ListMenu(r.Context())
	if err != nil {
		return err
	}

	return response(w, data(menu))
}

func (h MenuHandler) HandleFindMenu(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "menuID")
	if err != nil {
		return httputil.ErrBadRequest("ID Menu tidak valid!")
	}

	menu, err := h.menuUC.FindMenu(r.Context(), id)
	if err != nil {
		return err
	}

	return response(w, data(menu))
}

func (h MenuHandler) HandleDeleteKategori(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "categoryID")
	if err != nil {
		return httputil.ErrBadRequest("ID Kategori tidak valid!")
	}

	if err := h.menuUC.RemoveKategori(r.Context(), id); err != nil {
		return err
	}

	return response(w, message("Kategori menu berhasil dihapus!."))
}

func (h MenuHandler) HandleDeleteMenu(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "menuID")
	if err != nil {
		return httputil.ErrBadRequest("ID Menu tidak valid!")
	}

	if err := h.menuUC.RemoveMenu(r.Context(), id); err != nil {
		return err
	}

	return response(w, message("Menu berhasil dihapus!."))
}

func (h MenuHandler) HandleUpdateKategori(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "categoryID")
	if err != nil {
		return httputil.ErrBadRequest("ID Kategori tidak valid!")
	}

	reqBody, err := BindAndValidate[model.UpdateKategoriMenuRequest](r)
	if err != nil {
		return err
	}
	reqBody.ID = id

	if err := h.menuUC.UpdateKategori(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, message("Kategori menu berhasil di update!."))
}

func (h MenuHandler) HandleUpdateMenu(w http.ResponseWriter, r *http.Request) error {
	id, err := httputil.ParamInt(r, "menuID")
	if err != nil {
		return httputil.ErrBadRequest("ID Menu tidak valid!")
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

	// check if new image is included
	updateImage := true
	_, _, err = r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			updateImage = false
		} else {
			return err
		}
	}

	if updateImage {
		imageName, err := httputil.UploadHandler(r, "image")
		menuData.ImageName = imageName
		if err != nil {
			return err
		}
	}

	if err := h.menuUC.UpdateMenu(r.Context(), menuData); err != nil {
		return err
	}

	return response(w, message("Menu berhasil di update!."))
}

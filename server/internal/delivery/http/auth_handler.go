package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/usecase"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

type AuthHandler struct {
	userUC usecase.AuthUseCase
	cfg    *config.Config
}

func NewAuthHandler(cfg *config.Config, us usecase.AuthUseCase) AuthHandler {
	return AuthHandler{
		userUC: us,
		cfg:    cfg,
	}
}

func (h AuthHandler) HandleRegister(w http.ResponseWriter, r *http.Request) error {
	var reqBody model.RegisterRequest

	if err := httputil.BindJSON(r, &reqBody); err != nil {
		return err
	}

	if err := reqBody.ValidateRequest(); err != nil {
		return err
	}

	if err := h.userUC.Register(r.Context(), reqBody); err != nil {
		return err
	}

	return response(w, status(201), message("Registration successful."))
}

func (h AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	var reqBody model.LoginRequest
	web := strings.ToUpper(r.URL.Query().Get("web")) == "TRUE"

	if err := httputil.BindJSON(r, &reqBody); err != nil {
		return err
	}

	if err := reqBody.ValidateRequest(); err != nil {
		return err
	}

	tokenString, err := h.userUC.Login(r.Context(), reqBody, web)
	if err != nil {
		return err
	}

	tokenCookie := &http.Cookie{
		Name:     h.cfg.Token.AccessCookieName,
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(h.cfg.Token.Expiry * int(time.Hour))),
		HttpOnly: h.cfg.IsProduction(),
	}

	http.SetCookie(w, tokenCookie)

	d := map[string]string{
		"token": tokenString,
	}

	return response(w, data(d), message("Login successful."))
}

func (h AuthHandler) CurrentUser(w http.ResponseWriter, r *http.Request) error {
	tokenClaims := httputil.GetAuthCtx(r.Context())

	return response(w, data(tokenClaims["user"]))
}

func (h AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) error {
	tokenCookie := &http.Cookie{
		Name:     h.cfg.Token.AccessCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: h.cfg.IsProduction(),
	}

	http.SetCookie(w, tokenCookie)

	return response(w, message("Logout successful."))
}

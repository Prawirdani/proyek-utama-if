package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prawirdani/golang-restapi/internal/middleware"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var handlerFn = httputil.HandlerWrapper

func MapAuthRoutes(r chi.Router, h AuthHandler, mw middleware.MiddlewareManager) {
	r.Post("/auth/register", handlerFn(h.HandleRegister))
	r.Post("/auth/login", handlerFn(h.HandleLogin))
	r.Delete("/auth/logout", handlerFn(h.HandleLogout))
	r.With(mw.Authenticate).Get("/auth/current", handlerFn(h.CurrentUser))
}

func MapMenuRoutes(r chi.Router, h MenuHandler, mw middleware.MiddlewareManager) {
	r.With(mw.Authenticate).Group(func(subR chi.Router) {
		subR.Post("/menus/categories", handlerFn(h.HandleCreateKategori))
		subR.Get("/menus/categories", handlerFn(h.HandleListKategori))
		subR.Put("/menus/categories/{categoryID}", handlerFn(h.HandleUpdateKategori))
		subR.Delete("/menus/categories/{categoryID}", handlerFn(h.HandleDeleteKategori))

		subR.Get("/menus", handlerFn(h.HandleListMenu))
		subR.Get("/menus/{menuID}", handlerFn(h.HandleFindMenu))
		subR.Put("/menus/{menuID}", handlerFn(h.HandleUpdateMenu))
		subR.Delete("/menus/{menuID}", handlerFn(h.HandleDeleteMenu))
		subR.Post("/menus", handlerFn(h.HandleCreateMenu))
	})
}

func MapMejaRoutes(r chi.Router, h MejaHandler, mw middleware.MiddlewareManager) {
	r.With(mw.Authenticate).Group(func(subR chi.Router) {
		subR.Post("/tables", handlerFn(h.HandleCreateMeja))
		subR.Get("/tables", handlerFn(h.HandleListMeja))
		subR.Get("/tables/{tableID}", handlerFn(h.HandleFindMeja))
		subR.Put("/tables/{tableID}", handlerFn(h.HandleUpdateMeja))
		subR.Delete("/tables/{tableID}", handlerFn(h.HandleDeleteMeja))
	})
}

func MapPesananRoutes(r chi.Router, h PesananHandler, mw middleware.MiddlewareManager) {
	r.With(mw.Authenticate).Group(func(subR chi.Router) {
		subR.Get("/orders", handlerFn(h.HandleListPesanan))
		subR.Post("/orders/dinein", handlerFn(h.HandlePesananDineIn))
		subR.Post("/orders/takeaway", handlerFn(h.HandlePesananTakeAway))
		subR.Get("/orders/{pesananID}", handlerFn(h.HandleFindPesanan))
		subR.Put("/orders/{pesananID}/cancel", handlerFn(h.HandleBatalkanPesanan))
		subR.Put("/orders/{pesananID}/add-menu", handlerFn(h.HandleAddMenu))
		subR.Delete("/orders/{pesananID}/{detailID}", handlerFn(h.HandleRemoveMenu))
	})
}

func MapPembayaranRoutes(r chi.Router, h PembayaranHandler, mw middleware.MiddlewareManager) {
	r.With(mw.Authenticate).Group(func(subR chi.Router) {
		subR.Post("/payments/methods", handlerFn(h.HandleCreateMetodePembayaran))
		subR.Get("/payments/methods", handlerFn(h.HandleListMetodePembayaran))
		subR.Get("/payments/methods/{metodePembayaranID}", handlerFn(h.HandleFindMetodePembayaran))
		subR.Put("/payments/methods/{metodePembayaranID}", handlerFn(h.HandleUpdateMetodePembayaran))
		subR.Delete("/payments/methods/{metodePembayaranID}", handlerFn(h.HandleDeleteMetodePembayaran))
	})
}

// images file server, serving images from uploads directory
func ImagesFS(r chi.Router) {
	uploadsDir := "./uploads"
	r.Handle("/images/*", http.StripPrefix("/api/images/", http.FileServer(http.Dir(uploadsDir))))
}

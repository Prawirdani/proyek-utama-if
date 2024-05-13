package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/prawirdani/golang-restapi/internal/delivery/http"
	"github.com/prawirdani/golang-restapi/internal/middleware"
	"github.com/prawirdani/golang-restapi/internal/repository"
	"github.com/prawirdani/golang-restapi/internal/usecase"
)

// Init & Injects all dependencies.
func (s Server) bootstrap() {
	// Setup Repos
	userRepository := repository.NewUserRepository(s.pg, "users")
	menuRepository := repository.NewMenuRepository(s.pg, s.cfg)
	mejaRepository := repository.NewMejaRepository(s.pg, s.cfg)
	pesananRepository := repository.NewPesananRepository(s.pg, s.cfg)
	pembayaranRepository := repository.NewPembayaranRepository(s.pg, s.cfg)

	// Setup Usecases
	authUC := usecase.NewAuthUseCase(s.cfg, userRepository)
	menuUC := usecase.NewMenuUsecase(menuRepository, s.cfg)
	mejaUC := usecase.NewMejaUseCase(mejaRepository, s.cfg)
	pesananUC := usecase.NewPesananUseCase(s.cfg, menuRepository, mejaRepository, pesananRepository)
	pembayaranUC := usecase.NewPembayaranUsecase(s.cfg, pembayaranRepository, pesananRepository, mejaRepository)

	// Setup Handlers
	authHandler := http.NewAuthHandler(s.cfg, authUC)
	menuHandler := http.NewMenuHandler(s.cfg, menuUC)
	mejaHandler := http.NewMejaHandler(s.cfg, mejaUC)
	pembayaranHandler := http.NewPembayaranHandler(pembayaranUC, s.cfg)
	pesananHandler := http.NewPesananHandler(s.cfg, pesananUC)

	middlewares := middleware.NewMiddlewareManager(s.cfg)

	s.router.Route("/api", func(r chi.Router) {
		http.ImagesFS(r)
	})

	s.router.Route("/api/v1", func(v1 chi.Router) {
		http.MapAuthRoutes(v1, authHandler, middlewares)
		http.MapMenuRoutes(v1, menuHandler, middlewares)
		http.MapMejaRoutes(v1, mejaHandler, middlewares)
		http.MapPesananRoutes(v1, pesananHandler, middlewares)
		http.MapPembayaranRoutes(v1, pembayaranHandler, middlewares)
	})
}

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
	userRepository := repository.NewUserRepository(s.pg)
	menuRepository := repository.NewMenuRepository(s.pg, s.cfg)
	mejaRepository := repository.NewMejaRepository(s.pg, s.cfg)
	pesananRepository := repository.NewPesananRepository(s.pg, s.cfg)
	pembayaranRepository := repository.NewPembayaranRepository(s.pg, s.cfg)

	// Setup Usecases
	userUC := usecase.NewUserUseCase(s.cfg, userRepository)
	authUC := usecase.NewAuthUseCase(s.cfg, userRepository)
	menuUC := usecase.NewMenuUsecase(s.cfg, menuRepository)
	mejaUC := usecase.NewMejaUseCase(s.cfg, mejaRepository)
	pesananUC := usecase.NewPesananUseCase(s.cfg, menuRepository, mejaRepository, pesananRepository)
	pembayaranUC := usecase.NewPembayaranUsecase(s.cfg, pembayaranRepository, pesananRepository)

	// Setup Handlers
	userHandler := http.NewUserHandler(s.cfg, userUC)
	authHandler := http.NewAuthHandler(s.cfg, authUC)
	menuHandler := http.NewMenuHandler(s.cfg, menuUC)
	mejaHandler := http.NewMejaHandler(s.cfg, mejaUC)
	pembayaranHandler := http.NewPembayaranHandler(s.cfg, pembayaranUC)
	pesananHandler := http.NewPesananHandler(s.cfg, pesananUC)

	middlewares := middleware.NewMiddlewareManager(s.cfg)

	http.RegisterClientApp(s.router)

	s.router.Route("/api", func(r chi.Router) {
		http.ImagesFS(r)
	})

	s.router.Route("/api/v1", func(v1 chi.Router) {
		http.MapUserRoutes(v1, userHandler, middlewares)
		http.MapAuthRoutes(v1, authHandler, middlewares)
		http.MapMenuRoutes(v1, menuHandler, middlewares)
		http.MapMejaRoutes(v1, mejaHandler, middlewares)
		http.MapPesananRoutes(v1, pesananHandler, middlewares)
		http.MapPembayaranRoutes(v1, pembayaranHandler, middlewares)
	})
}

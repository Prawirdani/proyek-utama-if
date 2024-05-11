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
	paymentRepository := repository.NewPaymentRepository(s.pg, s.cfg)

	// Setup Usecases
	authUC := usecase.NewAuthUseCase(s.cfg, userRepository)
	menuUC := usecase.NewMenuUsecase(menuRepository, s.cfg)
	mejaUC := usecase.NewMejaUseCase(mejaRepository, s.cfg)
	paymentUC := usecase.NewPaymentUsecase(paymentRepository, s.cfg)

	// Setup Handlers
	authHandler := http.NewAuthHandler(s.cfg, authUC)
	menuHandler := http.NewMenuHandler(s.cfg, menuUC)
	mejaHandler := http.NewMejaHandler(s.cfg, mejaUC)
	paymentHandler := http.NewPaymentHandler(paymentUC, s.cfg)

	middlewares := middleware.NewMiddlewareManager(s.cfg)

	s.router.Route("/api", func(r chi.Router) {
		http.ImagesFS(r)
	})

	s.router.Route("/api/v1", func(v1 chi.Router) {
		http.MapAuthRoutes(v1, authHandler, middlewares)
		http.MapMenuRoutes(v1, menuHandler, middlewares)
		http.MapMejaRoutes(v1, mejaHandler, middlewares)
		http.MapPaymentRoutes(v1, paymentHandler, middlewares)
	})
}

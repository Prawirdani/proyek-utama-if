package usecase

import (
	"context"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type PesananUseCase interface {
	CreateDineIn(ctx context.Context, request model.PesananDineInRequest) error
	CreateTakeAway(ctx context.Context, request model.PesananTakeAwayRequest) error
	// FinishPesanan()
	// CancelPesanan()
	// AddMenuToPesanan()
	// RemoveMenuFromPesanan()
	// SetCatatan()
}

type pesananUseCase struct {
	cfg         *config.Config
	menuRepo    repository.MenuRepository
	mejaRepo    repository.MejaRepository
	pesananRepo repository.PesananRepository
}

func NewPesananUseCase(cfg *config.Config, menuRepo repository.MenuRepository, mejaRepo repository.MejaRepository, pesananRepo repository.PesananRepository) pesananUseCase {
	return pesananUseCase{
		cfg:         cfg,
		menuRepo:    menuRepo,
		mejaRepo:    mejaRepo,
		pesananRepo: pesananRepo,
	}
}

func (pu pesananUseCase) CreateDineIn(ctx context.Context, request model.PesananDineInRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout)*time.Second)
	defer cancel()

	// Retrieve meja then check is it Tersedia
	meja, err := pu.mejaRepo.SelectWhere(ctxWT, "id", request.MejaID)
	if err != nil {
		return err
	}

	// Return error if not tersedia
	if !meja.Tersedia() {
		return entity.ErrorMejaTidakTersedia
	}

	// Create new PesananDineIn
	pesanan := entity.NewPesananDineIn(request)

	// Retrieve menus & assign to Pesanan Detail
	if err := menusToDetails(ctxWT, pu.menuRepo, &pesanan, request.Menu...); err != nil {
		return err
	}

	if err := pu.pesananRepo.Insert(ctxWT, pesanan); err != nil {
		return err
	}

	return nil
}

func (pu pesananUseCase) CreateTakeAway(ctx context.Context, request model.PesananTakeAwayRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout)*time.Second)
	defer cancel()

	pesanan := entity.NewPesananTakeAway(request)

	if err := menusToDetails(ctxWT, pu.menuRepo, &pesanan, request.Menu...); err != nil {
		return err
	}

	if err := pu.pesananRepo.Insert(ctxWT, pesanan); err != nil {
		return err
	}

	return nil
}

// TODO: Should batch select rather that retrieve one by one
func menusToDetails(ctx context.Context, repo repository.MenuRepository, pesanan *entity.Pesanan, dr ...model.PesananMenuRequest) error {
	for i := 0; i < len(dr); i++ {
		d := dr[i]

		menu, err := repo.SelectWhere(ctx, "id", d.MenuID)
		if err != nil {
			return err
		}

		detail, err := entity.NewDetailPesanan(*menu, d.Kuantitas)
		if err != nil {
			return err
		}

		// Add detail to Pesanan
		pesanan.AddDetail(detail)
	}
	return nil
}

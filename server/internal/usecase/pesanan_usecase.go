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
	CreateDineIn(ctx context.Context, request model.PesananDineInRequest) (*int, error)
	CreateTakeAway(ctx context.Context, request model.PesananTakeAwayRequest) (*int, error)
	ListPesanan(ctx context.Context) ([]entity.Pesanan, error)
	FindPesanan(ctx context.Context, pesananID int) (*entity.Pesanan, error)
	FindPesananWithQuery(ctx context.Context, query *model.Query) (*entity.Pesanan, error)
	AddMenuToPesanan(ctx context.Context, pesananID int, request model.PesananMenuRequest) error
	RemoveMenuFromPesanan(ctx context.Context, pesananID int, detailID int) error
	BatalkanPesanan(ctx context.Context, pesananID int) error
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

func (pu pesananUseCase) FindPesananWithQuery(ctx context.Context, query *model.Query) (*entity.Pesanan, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout)*time.Second)
	defer cancel()

	return pu.pesananRepo.SelectQuery(ctxWT, query)
}

func (pu pesananUseCase) CreateDineIn(ctx context.Context, request model.PesananDineInRequest) (*int, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout)*time.Second)
	defer cancel()

	// Retrieve meja then check is it Tersedia
	meja, err := pu.mejaRepo.SelectWhere(ctxWT, "id", request.MejaID)
	if err != nil {
		return nil, err
	}

	// Create new PesananDineIn
	pesanan, err := entity.NewPesananDineIn(request, meja)
	if err != nil {
		return nil, err
	}

	// Retrieve menus & assign to Pesanan Detail
	if err := menusToDetails(ctxWT, pu.menuRepo, &pesanan, request.Menu...); err != nil {
		return nil, err
	}

	return pu.pesananRepo.Insert(ctxWT, pesanan)
}

func (pu pesananUseCase) CreateTakeAway(ctx context.Context, request model.PesananTakeAwayRequest) (*int, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout)*time.Second)
	defer cancel()

	pesanan := entity.NewPesananTakeAway(request)
	if err := menusToDetails(ctxWT, pu.menuRepo, &pesanan, request.Menu...); err != nil {
		return nil, err
	}

	return pu.pesananRepo.Insert(ctxWT, pesanan)
}

func (pu pesananUseCase) ListPesanan(ctx context.Context) ([]entity.Pesanan, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout)*time.Second)
	defer cancel()

	return pu.pesananRepo.Select(ctxWT)
}

func (pu pesananUseCase) FindPesanan(ctx context.Context, pesananID int) (*entity.Pesanan, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout)*time.Second)
	defer cancel()
	return pu.pesananRepo.SelectWhere(ctxWT, "p.id", pesananID)
}

func (pu pesananUseCase) AddMenuToPesanan(ctx context.Context, pesananID int, request model.PesananMenuRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	pesanan, err := pu.pesananRepo.SelectWhere(ctxWT, "p.id", pesananID)
	if err != nil {
		return err
	}

	err = menusToDetails(ctxWT, pu.menuRepo, pesanan, request)
	if err != nil {
		return err
	}

	err = pu.pesananRepo.Update(ctxWT, *pesanan)

	return err
}

func (pu pesananUseCase) BatalkanPesanan(ctx context.Context, pesananID int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	pesanan, err := pu.pesananRepo.SelectWhere(ctxWT, "p.id", pesananID)
	if err != nil {
		return err
	}

	err = pesanan.Batalkan()
	if err != nil {
		return err
	}

	err = pu.pesananRepo.Update(ctxWT, *pesanan)
	return err
}

func (pu pesananUseCase) RemoveMenuFromPesanan(ctx context.Context, pesananID int, detailID int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(pu.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	pesanan, err := pu.pesananRepo.SelectWhere(ctxWT, "p.id", pesananID)
	if err != nil {
		return err
	}

	err = pesanan.RemoveDetail(detailID)
	if err != nil {
		return err
	}

	err = pu.pesananRepo.DeleteDetail(ctxWT, *pesanan, detailID)
	if err != nil {
		return err
	}
	return err
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

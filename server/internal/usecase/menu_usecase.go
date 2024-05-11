package usecase

import (
	"context"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type MenuUsecase interface {
	// Kategori Contract
	CreateKategori(ctx context.Context, request model.CreateKategoriMenuRequest) error
	ListKategori(ctx context.Context) ([]entity.KategoriMenu, error)
	UpdateKategori(ctx context.Context, request model.UpdateKategoriMenuRequest) error
	RemoveKategori(ctx context.Context, id int) error
	// Menu Contract
	CreateMenu(ctx context.Context, request model.CreateMenuRequest) error
	ListMenu(ctx context.Context) ([]entity.Menu, error)
	FindMenu(ctx context.Context, id int) (*entity.Menu, error)
	UpdateMenu(ctx context.Context, request model.UpdateMenuRequest) error
	RemoveMenu(ctx context.Context, id int) error
}

type menuUsecase struct {
	menuRepo repository.MenuRepository
	cfg      *config.Config
}

func NewMenuUsecase(menuRepo repository.MenuRepository, cfg *config.Config) MenuUsecase {
	return &menuUsecase{
		menuRepo: menuRepo,
		cfg:      cfg,
	}
}

func (us menuUsecase) CreateKategori(ctx context.Context, request model.CreateKategoriMenuRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	newKategori := entity.NewKategoriMenu(request)

	return us.menuRepo.InsertKategori(ctxWT, newKategori)
}

func (us menuUsecase) ListKategori(ctx context.Context) ([]entity.KategoriMenu, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	return us.menuRepo.SelectKategori(ctxWT)
}

func (us menuUsecase) UpdateKategori(ctx context.Context, request model.UpdateKategoriMenuRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	kategori, err := us.menuRepo.SelectKategoriWhere(ctxWT, "id", request.ID)
	if err != nil {
		return err
	}
	kategori.Assign(request)

	return us.menuRepo.UpdateKategori(ctxWT, *kategori)
}

func (us menuUsecase) RemoveKategori(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	kategori, err := us.menuRepo.SelectKategoriWhere(ctxWT, "id", id)
	if err != nil {
		return err
	}
	kategori.SetDeletedAt()

	return us.menuRepo.UpdateKategori(ctxWT, *kategori)
}

func (us menuUsecase) CreateMenu(ctx context.Context, request model.CreateMenuRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	_, err := us.menuRepo.SelectKategoriWhere(ctxWT, "id", request.KategoriId)
	if err != nil {
		return err
	}

	newMenu := entity.NewMenu(request)
	return us.menuRepo.Insert(ctxWT, newMenu)
}

func (us menuUsecase) ListMenu(ctx context.Context) ([]entity.Menu, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	return us.menuRepo.Select(ctxWT)
}

func (us menuUsecase) FindMenu(ctx context.Context, id int) (*entity.Menu, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	return us.menuRepo.SelectWhere(ctxWT, "id", id)
}

func (us menuUsecase) UpdateMenu(ctx context.Context, request model.UpdateMenuRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()

	menu, err := us.FindMenu(ctxWT, request.ID)
	if err != nil {
		return err
	}

	menu.Assign(request)

	return us.menuRepo.Update(ctxWT, *menu)
}

func (us menuUsecase) RemoveMenu(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(us.cfg.Context.Timeout*int(time.Second)))
	defer cancel()
	menu, err := us.FindMenu(ctxWT, id)

	if err != nil {
		return err
	}
	menu.SetDeletedAt()

	return us.menuRepo.Update(ctxWT, *menu)
}

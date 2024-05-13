package usecase

import (
	"context"
	"time"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type PembayaranUseCase interface {
	CreateMetodePembayaran(ctx context.Context, request model.CreateMetodePembayaranRequest) error
	ListMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error)
	FindMetodePembayaran(ctx context.Context, id int) (*entity.MetodePembayaran, error)
	UpdateMetodePembayaran(ctx context.Context, request model.UpdateMetodePembayaranRequest) error
	RemoveMetodePembayaran(ctx context.Context, id int) error
	BayarPesanan(ctx context.Context, request model.PembayaranRequest) error
}

type pembayaranUsecase struct {
	pembayaranRepo repository.PembayaranRepository
	pesananRepo    repository.PesananRepository
	mejaRepo       repository.MejaRepository
	cfg            *config.Config
}

func NewPembayaranUsecase(
	cfg *config.Config,
	pembayaranRepo repository.PembayaranRepository,
	pesananRepo repository.PesananRepository,
	mejaRepo repository.MejaRepository,
) pembayaranUsecase {
	return pembayaranUsecase{
		cfg:            cfg,
		pembayaranRepo: pembayaranRepo,
		pesananRepo:    pesananRepo,
		mejaRepo:       mejaRepo,
	}
}

func (u pembayaranUsecase) CreateMetodePembayaran(ctx context.Context, request model.CreateMetodePembayaranRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := entity.NewMetodePembayaran(request)
	if err != nil {
		return err
	}

	err = u.pembayaranRepo.InsertMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

func (u pembayaranUsecase) ListMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mps, err := u.pembayaranRepo.SelectMetodePembayaran(ctxWT)
	if err != nil {
		return nil, err
	}

	return mps, nil
}

func (u pembayaranUsecase) FindMetodePembayaran(ctx context.Context, id int) (*entity.MetodePembayaran, error) {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.pembayaranRepo.SelectMetodePembayaranWhere(ctxWT, "id", id)
	if err != nil {
		return nil, err
	}

	return mp, nil
}

func (u pembayaranUsecase) UpdateMetodePembayaran(ctx context.Context, request model.UpdateMetodePembayaranRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.FindMetodePembayaran(ctxWT, request.ID)
	if err != nil {
		return err
	}

	err = mp.Assign(request)
	if err != nil {
		return err
	}

	err = u.pembayaranRepo.UpdateMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

func (u pembayaranUsecase) RemoveMetodePembayaran(ctx context.Context, id int) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	mp, err := u.FindMetodePembayaran(ctxWT, id)
	if err != nil {
		return err
	}

	mp.SetDeletedAt()
	err = u.pembayaranRepo.UpdateMetodePembayaran(ctxWT, *mp)
	if err != nil {
		return err
	}

	return nil
}

func (u pembayaranUsecase) BayarPesanan(ctx context.Context, request model.PembayaranRequest) error {
	ctxWT, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout)*time.Second)
	defer cancel()

	// Find Pesanan
	pesanan, err := u.pesananRepo.SelectWhere(ctxWT, "p.id", request.PesananId)
	if err != nil {
		return err
	}
	// Set pesanan to selesai
	if err := pesanan.Selesaikan(); err != nil {
		return err
	}

	// Find Metode Pembayaran
	metodePembayaran, err := u.pembayaranRepo.SelectMetodePembayaranWhere(ctxWT, "id", request.MetodePembayaranId)
	if err != nil {
		return err
	}

	// Create Pembayaran
	pembayaran := entity.NewPembayaran(*pesanan, *metodePembayaran)

	return u.pembayaranRepo.CreatePembayaran(ctxWT, *pesanan, *pembayaran)
}

package repository

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
)

type PesananRepository interface {
	Insert(ctx context.Context, p entity.Pesanan) error
}

type pesananRepository struct {
	db  *pgxpool.Pool
	cfg *config.Config
}

func NewPesananRepository(db *pgxpool.Pool, cfg *config.Config) pesananRepository {
	return pesananRepository{
		db:  db,
		cfg: cfg,
	}
}

func (pr pesananRepository) Insert(ctx context.Context, p entity.Pesanan) error {
	query := `INSERT INTO pesanan (nama_pelanggan, kasir_id, meja_id, tipe_pesanan, status_pesanan, catatan, total, waktu_pesanan) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	tx, err := pr.db.Begin(ctx)
	if err != nil {
		slog.Error("PesananRepository.Insert.Begin", slog.Any("error", err))
		return err
	}
	defer tx.Rollback(ctx)

	var pesananID int

	mejaID := func() *int {
		if p.Meja != nil {
			return &p.Meja.ID
		}
		return nil
	}()

	err = tx.QueryRow(ctx, query,
		p.NamaPelanggan,
		p.Kasir.ID,
		mejaID,
		p.TipePesanan,
		p.StatusPesanan,
		p.Catatan,
		p.Total,
		p.WaktuPesanan,
	).Scan(&pesananID)

	if err != nil {
		slog.Error("PesananRepository.Insert.Exec.Pesanan", slog.Any("error", err))
		return err
	}

	// Insert pesanan details
	_, err = tx.CopyFrom(
		ctx,
		pgx.Identifier{"detail_pesanan"},
		[]string{"pesanan_id", "menu_id", "harga", "kuantitas", "subtotal"},
		pgx.CopyFromSlice(len(p.Detail), func(i int) ([]interface{}, error) {
			d := p.Detail[i]
			return []interface{}{pesananID, d.Menu.ID, d.Menu.Harga, d.Kuantitas, d.Subtotal}, nil
		}),
	)
	if err != nil {
		slog.Error("PesananRepository.Insert.Exec.DetailPesanan", slog.Any("error", err))
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

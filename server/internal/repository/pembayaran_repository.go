package repository

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
)

type PembayaranRepository interface {
	InsertMetodePembayaran(ctx context.Context, mp entity.MetodePembayaran) error
	SelectMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error)
	SelectMetodePembayaranWhere(ctx context.Context, field string, searchVal any) (*entity.MetodePembayaran, error)
	UpdateMetodePembayaran(ctx context.Context, mp entity.MetodePembayaran) error
	InsertPembayaran(ctx context.Context, pesanan entity.Pesanan, pembayaran entity.Pembayaran) (*entity.Receipt, error)
}

type pembayaranRepository struct {
	db  *pgxpool.Pool
	cfg *config.Config
}

func NewPembayaranRepository(db *pgxpool.Pool, cfg *config.Config) pembayaranRepository {
	return pembayaranRepository{
		db:  db,
		cfg: cfg,
	}
}

func (r pembayaranRepository) InsertMetodePembayaran(ctx context.Context, mp entity.MetodePembayaran) error {
	query := "INSERT INTO metode_pembayaran (tipe_pembayaran, metode, deskripsi) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(ctx, query, mp.TipePembayaran, mp.Metode, mp.Deskripsi)
	if err != nil {
		// Duplicate unique error
		if strings.Contains(err.Error(), "23505") {
			return entity.ErrorMetodePembayaranDuplicate
		}
		return err
	}

	return nil
}

func (r pembayaranRepository) SelectMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error) {
	query := "SELECT id, tipe_pembayaran, metode, deskripsi, deleted_at FROM metode_pembayaran WHERE deleted_at IS NULL"

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metodePembayaran []entity.MetodePembayaran
	for rows.Next() {
		var mp entity.MetodePembayaran
		err := mp.ScanRow(rows)
		if err != nil {
			return nil, err
		}
		metodePembayaran = append(metodePembayaran, mp)
	}

	return metodePembayaran, nil
}

func (r pembayaranRepository) SelectMetodePembayaranWhere(ctx context.Context, field string, searchVal any) (*entity.MetodePembayaran, error) {
	query := fmt.Sprintf("SELECT id, tipe_pembayaran, metode, deskripsi, deleted_at FROM metode_pembayaran WHERE %s=$1 AND deleted_at IS NULL", field)

	row := r.db.QueryRow(ctx, query, searchVal)

	var mp entity.MetodePembayaran
	err := mp.ScanRow(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrorMetodePembayaranNotFound
		}
		return nil, err
	}

	return &mp, nil
}

func (r pembayaranRepository) UpdateMetodePembayaran(ctx context.Context, mp entity.MetodePembayaran) error {
	query := "UPDATE metode_pembayaran SET tipe_pembayaran=$1, metode=$2, deskripsi=$3, deleted_at=$4 WHERE id=$5"
	_, err := r.db.Exec(ctx, query, mp.TipePembayaran, mp.Metode, mp.Deskripsi, mp.DeletedAt, mp.ID)
	if err != nil {
		// Duplicate unique error
		if strings.Contains(err.Error(), "23505") {
			return entity.ErrorMetodePembayaranDuplicate
		}
		return err
	}

	return nil
}

func (r pembayaranRepository) InsertPembayaran(ctx context.Context, pesanan entity.Pesanan, pembayaran entity.Pembayaran) (*entity.Receipt, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		slog.Error("PesananRepository.CreatePembayaran.Begin", slog.Any("error", err))
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `
	WITH inserted AS (
		INSERT INTO pembayaran (pesanan_id, metode_pembayaran_id, jumlah, waktu_pembayaran)
	    VALUES ($1, $2, $3, $4)
	    RETURNING id, pesanan_id, metode_pembayaran_id, jumlah, waktu_pembayaran
	)
	SELECT 
		i.id, i.jumlah, i.waktu_pembayaran,
	    mp.id, mp.tipe_pembayaran, mp.metode, mp.deskripsi
	FROM inserted AS i
	JOIN metode_pembayaran AS mp ON i.metode_pembayaran_id = mp.id
	`
	var p entity.Pembayaran
	if err := tx.QueryRow(ctx, query, pesanan.ID, pembayaran.Metode.ID, pembayaran.Jumlah, pembayaran.WaktuPembayaran).Scan(
		&p.Id,
		&p.Jumlah,
		&p.WaktuPembayaran,
		&p.Metode.ID,
		&p.Metode.TipePembayaran,
		&p.Metode.Metode,
		&p.Metode.Deskripsi,
	); err != nil {
		slog.Error("PesananRepository.CreatePembayaran.Exec", slog.Any("error", err))
		return nil, err
	}

	if err := updatePesanan(ctx, tx, pesanan); err != nil {
		return nil, err
	}

	// Update meja if Pesanan Dine In
	if pesanan.IsDineIn() {
		if err := updateMeja(ctx, tx, *pesanan.Meja); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	receipt := entity.NewInvoice(pesanan, p)

	return receipt, nil
}

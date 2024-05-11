package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/valueobject"
)

type PembayaranRepository interface {
	InsertMetodePembayaran(ctx context.Context, mp valueobject.MetodePembayaran) error
	SelectMetodePembayaran(ctx context.Context) ([]valueobject.MetodePembayaran, error)
	SelectMetodePembayaranWhere(ctx context.Context, field string, searchVal any) (*valueobject.MetodePembayaran, error)
	UpdateMetodePembayaran(ctx context.Context, mp valueobject.MetodePembayaran) error
}

type pembayaranRepository struct {
	db  *pgxpool.Pool
	cfg *config.Config
}

func NewPembayaranRepository(db *pgxpool.Pool, cfg *config.Config) PembayaranRepository {
	return &pembayaranRepository{
		db:  db,
		cfg: cfg,
	}
}

func (r pembayaranRepository) InsertMetodePembayaran(ctx context.Context, mp valueobject.MetodePembayaran) error {
	query := "INSERT INTO metode_pembayaran (tipe_pembayaran, metode, deskripsi) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(ctx, query, mp.TipePembayaran, mp.Metode, mp.Deskripsi)
	if err != nil {
		// Duplicate unique error
		if strings.Contains(err.Error(), "23505") {
			return valueobject.ErrorMetodePembayaranDuplicate
		}
		return err
	}

	return nil
}

func (r pembayaranRepository) SelectMetodePembayaran(ctx context.Context) ([]valueobject.MetodePembayaran, error) {
	query := "SELECT id, tipe_pembayaran, metode, deskripsi, deleted_at FROM metode_pembayaran WHERE deleted_at IS NULL"

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metodePembayaran []valueobject.MetodePembayaran
	for rows.Next() {
		var mp valueobject.MetodePembayaran
		err := mp.ScanRow(rows)
		if err != nil {
			return nil, err
		}
		metodePembayaran = append(metodePembayaran, mp)
	}

	return metodePembayaran, nil
}

func (r pembayaranRepository) SelectMetodePembayaranWhere(ctx context.Context, field string, searchVal any) (*valueobject.MetodePembayaran, error) {
	query := fmt.Sprintf("SELECT id, tipe_pembayaran, metode, deskripsi, deleted_at FROM metode_pembayaran WHERE %s=$1 AND deleted_at IS NULL", field)

	row := r.db.QueryRow(ctx, query, searchVal)

	var mp valueobject.MetodePembayaran
	err := mp.ScanRow(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, valueobject.ErrorMetodePembayaranNotFound
		}
		return nil, err
	}

	return &mp, nil
}

func (r pembayaranRepository) UpdateMetodePembayaran(ctx context.Context, mp valueobject.MetodePembayaran) error {
	query := "UPDATE metode_pembayaran SET tipe_pembayaran=$1, metode=$2, deskripsi=$3, deleted_at=$4 WHERE id=$5"
	_, err := r.db.Exec(ctx, query, mp.TipePembayaran, mp.Metode, mp.Deskripsi, mp.DeletedAt, mp.ID)
	if err != nil {
		// Duplicate unique error
		if strings.Contains(err.Error(), "23505") {
			return valueobject.ErrorMetodePembayaranDuplicate
		}
		return err
	}

	return nil
}

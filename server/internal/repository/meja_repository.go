package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
)

type MejaRepository interface {
	Insert(ctx context.Context, meja entity.Meja) error
	Select(ctx context.Context) ([]entity.Meja, error)
	SelectWhere(ctx context.Context, field string, searchVal any) (*entity.Meja, error)
	Update(ctx context.Context, meja entity.Meja) error
}

type mejaRepository struct {
	db  *pgxpool.Pool
	cfg *config.Config
}

func NewMejaRepository(db *pgxpool.Pool, cfg *config.Config) mejaRepository {
	return mejaRepository{
		db:  db,
		cfg: cfg,
	}
}

func (m mejaRepository) Insert(ctx context.Context, meja entity.Meja) error {
	query := "INSERT INTO meja (nomor, status) VALUES ($1, $2)"
	_, err := m.db.Exec(ctx, query, meja.Nomor, meja.Status)
	if err != nil {
		// Duplicate Unique Key error
		if strings.Contains(err.Error(), "23505") {
			return entity.ErrorMejaDuplicate
		}
		return err
	}
	return nil
}

func (m mejaRepository) Select(ctx context.Context) ([]entity.Meja, error) {
	query := "SELECT id, nomor, status, deleted_at FROM meja WHERE deleted_at IS NULL ORDER BY nomor ASC"
	rows, err := m.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	mejas := make([]entity.Meja, 0)
	for rows.Next() {
		var meja entity.Meja
		err := meja.ScanRow(rows)
		if err != nil {
			return nil, err
		}
		mejas = append(mejas, meja)
	}

	return mejas, nil
}

func (m mejaRepository) SelectWhere(ctx context.Context, field string, searchVal any) (*entity.Meja, error) {
	query := fmt.Sprintf("SELECT id, nomor, status, deleted_at FROM meja WHERE %s=$1 AND deleted_at IS NULL", field)

	var meja entity.Meja
	row := m.db.QueryRow(ctx, query, searchVal)

	err := meja.ScanRow(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrorMejaNotFound
		}
		return nil, nil
	}

	return &meja, nil
}

func (m mejaRepository) Update(ctx context.Context, meja entity.Meja) error {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := updateMeja(ctx, tx, meja); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func updateMeja(ctx context.Context, tx pgx.Tx, meja entity.Meja) error {
	query := "UPDATE meja SET nomor=$1, status=$2, deleted_at=$3 WHERE id=$4"
	_, err := tx.Exec(ctx, query, meja.Nomor, meja.Status, meja.DeletedAt, meja.ID)
	if err != nil {
		// Duplicate Unique Key error
		if strings.Contains(err.Error(), "23505") {
			return entity.ErrorMejaDuplicate
		}
		return err
	}

	return nil
}

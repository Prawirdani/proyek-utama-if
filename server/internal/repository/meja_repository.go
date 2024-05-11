package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
)

type MejaRepository interface {
	Insert(ctx context.Context, meja entity.Meja) error
	Select(ctx context.Context) ([]entity.Meja, error)
	SelectWhere(ctx context.Context, field string, searchVal any) (entity.Meja, error)
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
	return nil
}

func (m mejaRepository) Select(ctx context.Context) ([]entity.Meja, error) {
	return []entity.Meja{}, nil
}

func (m mejaRepository) SelectWhere(ctx context.Context, field string, searchVal any) (entity.Meja, error) {
	return entity.Meja{}, nil
}

func (m mejaRepository) Update(ctx context.Context, meja entity.Meja) error {
	return nil
}

package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
)

type PaymentRepository interface {
}

type paymentRepository struct {
	db  *pgxpool.Pool
	cfg *config.Config
}

func NewPaymentRepository(db *pgxpool.Pool, cfg *config.Config) PaymentRepository {
	return &paymentRepository{
		db:  db,
		cfg: cfg,
	}
}

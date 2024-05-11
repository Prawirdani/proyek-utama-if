package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
)

type MenuRepository interface {
	InsertKategori(ctx context.Context, m entity.KategoriMenu) error
	UpdateKategori(ctx context.Context, m entity.KategoriMenu) error
	SelectKategori(ctx context.Context) ([]entity.KategoriMenu, error)
	SelectKategoriWhere(ctx context.Context, field string, searchVal any) (*entity.KategoriMenu, error)
	Insert(ctx context.Context, m entity.Menu) error
	Update(ctx context.Context, m entity.Menu) error
	Select(ctx context.Context) ([]entity.Menu, error)
	SelectWhere(ctx context.Context, field string, searchVal any) (*entity.Menu, error)
}

type menuRepository struct {
	db  *pgxpool.Pool
	cfg *config.Config
}

func NewMenuRepository(pgpool *pgxpool.Pool, cfg *config.Config) menuRepository {
	return menuRepository{
		db:  pgpool,
		cfg: cfg,
	}
}

func (r menuRepository) InsertKategori(ctx context.Context, m entity.KategoriMenu) error {
	query := "INSERT INTO kategori_menu (nama) VALUES ($1)"
	_, err := r.db.Exec(ctx, query, m.Nama)
	if err != nil {
		return err
	}
	return nil
}

func (r menuRepository) UpdateKategori(ctx context.Context, m entity.KategoriMenu) error {
	query := "UPDATE kategori_menu SET nama=$1, deleted_at=$2 WHERE id=$3"
	_, err := r.db.Exec(ctx, query, m.Nama, m.DeletedAt, m.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r menuRepository) SelectKategori(ctx context.Context) ([]entity.KategoriMenu, error) {
	query := "SELECT * FROM kategori_menu WHERE deleted_at IS NULL"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kategories []entity.KategoriMenu
	for rows.Next() {
		var k entity.KategoriMenu
		err := k.ScanRow(rows)
		if err != nil {
			return nil, err
		}
		kategories = append(kategories, k)
	}
	return kategories, nil
}

func (r menuRepository) SelectKategoriWhere(ctx context.Context, field string, searchVal any) (*entity.KategoriMenu, error) {
	query := fmt.Sprintf("SELECT * FROM kategori_menu WHERE %s=$1 AND deleted_at IS NULL", field)
	row := r.db.QueryRow(ctx, query, searchVal)

	var kategori entity.KategoriMenu

	err := kategori.ScanRow(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrorKategoriNotFound
		}
		return nil, err
	}
	return &kategori, nil
}

func (r menuRepository) Insert(ctx context.Context, m entity.Menu) error {
	query := "INSERT INTO menus (nama, deskripsi, harga, kategori_id, url_foto) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(ctx, query, m.Nama, m.Deskripsi, m.Harga, m.Kategori.ID, m.Url)
	if err != nil {
		// Violate Foreign Key err by PG error code.
		if strings.Contains(err.Error(), "23503") {
			return entity.ErrorKategoriNotFound
		}
		return err
	}
	return nil
}

func (r menuRepository) Update(ctx context.Context, m entity.Menu) error {
	updatedAt := time.Now()
	query := "UPDATE menus SET nama=$1, deskripsi=$2, harga=$3, kategori_id=$4, url_foto=$5, deleted_at=$6, updated_at=$7 WHERE id=$8"
	_, err := r.db.Exec(ctx, query, m.Nama, m.Deskripsi, m.Harga, m.Kategori.ID, m.Url, m.DeletedAt, updatedAt, m.ID)
	if err != nil {
		// Violate Foreign Key err by PG error code.
		if strings.Contains(err.Error(), "23503") {
			return entity.ErrorKategoriNotFound
		}
		return err
	}
	return nil
}

func (r menuRepository) Select(ctx context.Context) ([]entity.Menu, error) {
	query := querySelectMenu + " WHERE m.deleted_at IS NULL"

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []entity.Menu
	for rows.Next() {
		var m entity.Menu

		if err := m.ScanRow(rows); err != nil {
			return nil, err
		}
		m.FormatURL(r.cfg)
		menus = append(menus, m)
	}
	return menus, nil
}

func (r menuRepository) SelectWhere(ctx context.Context, field string, searchVal any) (*entity.Menu, error) {
	query := querySelectMenu + fmt.Sprintf(" WHERE m.%s=$1 AND m.deleted_at IS NULL", field)
	row := r.db.QueryRow(ctx, query, searchVal)

	var menu entity.Menu
	if err := menu.ScanRow(row); err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrorMenuNotFound
		}
		return nil, err
	}
	return &menu, nil
}

var (
	querySelectMenu = `
	SELECT
		m.id, m.nama, m.deskripsi, m.harga, m.url_foto, m.deleted_at, m.created_at, m.updated_at,
		km.id, km.nama, km.deleted_at
	FROM menus AS m
		JOIN kategori_menu AS km ON m.kategori_id=km.id
	`
)

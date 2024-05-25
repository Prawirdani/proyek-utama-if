package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/internal/model"
)

type PesananRepository interface {
	Insert(ctx context.Context, p entity.Pesanan) error
	Select(ctx context.Context) ([]entity.Pesanan, error)
	SelectWhere(ctx context.Context, field string, searchVal any) (*entity.Pesanan, error)
	// Complex Search With Query Params
	SelectQuery(ctx context.Context, query *model.Query) (*entity.Pesanan, error)
	Update(ctx context.Context, pesanan entity.Pesanan) error
	DeleteDetail(ctx context.Context, pesanan entity.Pesanan, detailID int) error
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

func (pr pesananRepository) SelectQuery(ctx context.Context, query *model.Query) (*entity.Pesanan, error) {
	var pesanan entity.Pesanan
	baseQuery := querySelectPesanan
	queryStr := query.Build(baseQuery)

	row := pr.db.QueryRow(ctx, queryStr, query.StmtArgs...)
	err := pesanan.ScanRow(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrorPesananNotFound
		}
		return nil, err
	}

	details, err := pr.queryPesananDetails(ctx, pesanan.ID)
	pesanan.Detail = append(pesanan.Detail, details...)

	return &pesanan, nil
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

	err = pr.batchDetailInsert(ctx, tx, pesananID, p.Detail)
	if err != nil {
		slog.Error("PesananRepository.Insert.Exec.DetailPesanan", slog.Any("error", err))
		return err
	}

	// if pesanan is dine in update meja status to terisi
	if p.IsDineIn() {
		if err := updateMeja(ctx, tx, *p.Meja); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (pr pesananRepository) Select(ctx context.Context) ([]entity.Pesanan, error) {
	var ps []entity.Pesanan

	rows, err := pr.db.Query(ctx, querySelectPesanan)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := entity.Pesanan{
			Kasir: entity.User{},
		}
		err := p.ScanRow(rows)
		if err != nil {
			return nil, err
		}

		details, err := pr.queryPesananDetails(ctx, p.ID)
		if err != nil {
			return nil, err
		}

		p.Detail = append(p.Detail, details...)
		ps = append(ps, p)
	}
	return ps, nil
}

func (pr pesananRepository) SelectWhere(ctx context.Context, field string, searchVal any) (*entity.Pesanan, error) {
	var pesanan entity.Pesanan
	query := querySelectPesanan + fmt.Sprintf(" WHERE %s=$1", field)

	row := pr.db.QueryRow(ctx, query, searchVal)

	err := pesanan.ScanRow(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrorPesananNotFound
		}
		return nil, err
	}

	details, err := pr.queryPesananDetails(ctx, pesanan.ID)
	pesanan.Detail = append(pesanan.Detail, details...)

	return &pesanan, nil
}

func (pr pesananRepository) Update(ctx context.Context, pesanan entity.Pesanan) error {

	tx, err := pr.db.Begin(ctx)
	if err != nil {
		slog.Error("PesananRepository.Update.Begin", slog.Any("error", err))
		return err
	}
	defer tx.Rollback(ctx)

	err = updatePesanan(ctx, tx, pesanan)
	if err != nil {
		return err
	}

	// Seperating new added details and stored details, new added details via usecase will have ID=0
	newDetail, _ := func() (new []entity.DetailPesanan, stored []entity.DetailPesanan) {
		for _, d := range pesanan.Detail {
			if d.ID == 0 {
				new = append(new, d)
			} else {
				stored = append(stored, d)
			}
		}
		return new, stored
	}()

	// if there are new details, insert them
	if newDetail != nil {
		err = pr.batchDetailInsert(ctx, tx, pesanan.ID, newDetail)
		if err != nil {
			slog.Error("PesananRepository.Update.Insert.DetailPesanan", slog.Any("error", err))
			return err
		}
	}

	if pesanan.IsDineIn() {
		if err := updateMeja(ctx, tx, *pesanan.Meja); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (pr pesananRepository) DeleteDetail(ctx context.Context, pesanan entity.Pesanan, detailID int) error {
	tx, err := pr.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = updatePesanan(ctx, tx, pesanan)
	if err != nil {
		return err
	}

	query := `DELETE FROM detail_pesanan WHERE id=$1`
	_, err = tx.Exec(ctx, query, detailID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// updatePesanan is a helper function to update pesanan data
// shared across repositories
func updatePesanan(ctx context.Context, tx pgx.Tx, pesanan entity.Pesanan) error {
	query := `UPDATE pesanan 
		SET nama_pelanggan=$1, meja_id=$2, total=$3, tipe_pesanan=$4, status_pesanan=$5, catatan=$6
	WHERE id=$7
	`
	_, err := tx.Exec(ctx, query,
		pesanan.NamaPelanggan, pesanan.Meja.ID, pesanan.Total, pesanan.TipePesanan, pesanan.StatusPesanan, pesanan.Catatan, pesanan.ID,
	)
	return err
}

func (pr pesananRepository) batchDetailInsert(ctx context.Context, tx pgx.Tx, pesananID int, details []entity.DetailPesanan) error {
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"detail_pesanan"},
		[]string{"pesanan_id", "menu_id", "harga", "kuantitas", "subtotal"},
		pgx.CopyFromSlice(len(details), func(i int) ([]interface{}, error) {
			d := details[i]
			return []interface{}{pesananID, d.Menu.ID, d.Menu.Harga, d.Kuantitas, d.Subtotal}, nil
		}),
	)
	return err
}

func (pr pesananRepository) queryPesananDetails(ctx context.Context, pesananID int) ([]entity.DetailPesanan, error) {
	var details []entity.DetailPesanan
	rows, err := pr.db.Query(ctx, querySelectDetailPesanan, pesananID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		d := entity.DetailPesanan{
			Menu: entity.Menu{},
		}
		err := rows.Scan(
			&d.ID,
			&d.Menu.Harga,
			&d.Kuantitas,
			&d.Subtotal,
			&d.Menu.Nama,
		)
		if err != nil {
			return nil, err
		}
		details = append(details, d)
	}
	return details, nil
}

var (
	querySelectPesanan = `
	SELECT 
	p.id, p.nama_pelanggan, p.total, p.tipe_pesanan, p.status_pesanan, p.catatan, p.waktu_pesanan,
	k.nama,
	m.id, m.nomor, m.status
	FROM pesanan AS p
		JOIN users AS k ON p.kasir_id = k.id
		LEFT JOIN meja AS m ON p.meja_id = m.id
	`
	querySelectDetailPesanan = `
	SELECT
		dp.id, dp.harga, dp.kuantitas, dp.subtotal,
		m.nama AS nama_menu
	FROM detail_pesanan AS dp
		JOIN menus AS m ON dp.menu_id=m.id
	WHERE dp.pesanan_id=$1
	`
)

// SELECT
//     p.id,
//     p.nama_pelanggan,
//     p.total,
//     p.tipe_pesanan,
//     p.status_pesanan,
//     p.catatan,
//     p.waktu_pesanan,
//     k.nama AS kasir_nama,
//     mj.id AS meja_id,
//     mj.nomor AS meja_nomor,
//     dp.id AS detail_id,
//     dp.harga,
//     dp.kuantitas,
//     dp.subtotal,
//     m.nama AS menu_nama,
//     km.nama AS kategori_menu_nama
// FROM pesanan AS p
// JOIN users AS k ON p.kasir_id = k.id
// LEFT JOIN meja AS mj ON p.meja_id = mj.id
// LEFT JOIN detail_pesanan AS dp ON p.id = dp.pesanan_id
// LEFT JOIN menus AS m ON dp.menu_id = m.id
// LEFT JOIN kategori_menu AS km ON m.kategori_id = km.id;
//

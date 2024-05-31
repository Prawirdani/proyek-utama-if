package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/model"
)

type ReportRepository interface {
	DailyReport(context.Context) (model.Report, error)
}

type reportRepository struct {
	cfg *config.Config
	db  *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool, cfg *config.Config) reportRepository {
	return reportRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r reportRepository) DailyReport(ctx context.Context) (model.Report, error) {

	var report model.Report
	reportChan := make(chan model.Report)
	errChan := make(chan error, 2) // Buffered channel to avoid blocking

	go func() {
		pesananReport, err := r.queryPesananReport(ctx)
		if err != nil {
			errChan <- err
			return
		}
		report.PesananReport = pesananReport
		reportChan <- report // Send the updated report to the channel
	}()

	go func() {
		menuReport, err := r.queryMenuReport(ctx)
		if err != nil {
			errChan <- err
			return
		}
		report.MenuReport = menuReport
		reportChan <- report // Send the updated report to the channel
	}()

	for i := 0; i < 2; i++ {
		select {
		case report := <-reportChan:
			// Receive and ignore the report since we have already updated the main report
			_ = report
		case err := <-errChan:
			// Error received, return early
			return report, err
		}
	}

	return report, nil
}

func (r reportRepository) queryPesananReport(ctx context.Context) ([]model.PesananReport, error) {
	var report []model.PesananReport

	rows, err := r.db.Query(ctx, pesananReportQuery)
	if err != nil {
		return report, err
	}
	for rows.Next() {
		var p model.PesananReport
		err := rows.Scan(&p.PesananID, &p.NamaPelanggan, &p.TipePesanan, &p.WaktuPesanan, &p.Kasir, &p.Total, &p.MetodePembayaran)
		if err != nil {
			return report, err
		}
		report = append(report, p)
	}

	return report, nil
}

func (r reportRepository) queryMenuReport(ctx context.Context) ([]model.MenuReport, error) {
	var report []model.MenuReport

	rows, err := r.db.Query(ctx, menuReportQuery)
	if err != nil {
		return report, err
	}
	for rows.Next() {
		var m model.MenuReport
		err := rows.Scan(&m.MenuID, &m.NamaMenu, &m.Kategori, &m.JumlahTerjual, &m.TotalPendapatan)
		if err != nil {
			return report, err
		}
		report = append(report, m)
	}

	return report, nil
}

const (
	pesananReportQuery = `
	SELECT 
	p.id, p.nama_pelanggan, p.tipe_pesanan, p.waktu_pesanan,
	k.nama AS kasir,
	pm.jumlah,
	mp.metode
	FROM pesanan AS p
		JOIN users AS k ON p.kasir_id = k.id
		JOIN pembayaran AS pm ON pm.pesanan_id = p.id
		JOIN metode_pembayaran AS mp ON pm.metode_pembayaran_id = mp.id
	WHERE status_pesanan='Selesai';
	`

	menuReportQuery = `
	SELECT
		m.id AS menu_id,
		m.nama AS menu_name,
		km.nama AS kategori,
		SUM(dp.kuantitas) AS order_count,
		SUM(dp.subtotal) AS total_income
	FROM
		detail_pesanan AS dp
	JOIN
		menus AS m ON dp.menu_id = m.id
	JOIN
		kategori_menu AS km ON m.kategori_id = km.id
	JOIN
		pesanan AS p ON dp.pesanan_id = p.id
	WHERE p.status_pesanan = 'Selesai'
	GROUP BY
		m.id, m.nama, km.nama
	ORDER BY
		total_income DESC;
	`
)

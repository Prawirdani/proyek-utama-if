package usecase

import (
	"bytes"
	"context"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/repository"
)

type ReportUsecase interface {
	GenerateReport(context.Context) (*bytes.Buffer, error)
}

type reportUsecase struct {
	cfg        *config.Config
	reportRepo repository.ReportRepository
}

func NewReportUsecase(cfg *config.Config, reportRepo repository.ReportRepository) reportUsecase {
	return reportUsecase{
		cfg:        cfg,
		reportRepo: reportRepo,
	}
}

func (r reportUsecase) GenerateReport(ctx context.Context) (*bytes.Buffer, error) {
	report, err := r.reportRepo.DailyReport(ctx)
	if err != nil {
		return nil, err
	}

	return report.GenerateExcel()
}

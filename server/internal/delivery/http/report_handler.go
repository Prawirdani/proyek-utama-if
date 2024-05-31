package http

import (
	"fmt"
	"net/http"

	"github.com/prawirdani/golang-restapi/config"
	"github.com/prawirdani/golang-restapi/internal/usecase"
)

type ReportHandler struct {
	cfg           *config.Config
	reportUsecase usecase.ReportUsecase
}

func NewReportHandler(cfg *config.Config, reportUsecase usecase.ReportUsecase) ReportHandler {
	return ReportHandler{
		cfg:           cfg,
		reportUsecase: reportUsecase,
	}
}

func (h ReportHandler) GenerateReport(w http.ResponseWriter, r *http.Request) error {
	buf, err := h.reportUsecase.GenerateReport(r.Context())
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))

	// Write the buffer to the response writer
	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	return err
}

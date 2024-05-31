package model

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

type PesananReport struct {
	PesananID        int
	NamaPelanggan    string
	TipePesanan      string
	Kasir            string
	Total            int
	MetodePembayaran string
	WaktuPesanan     time.Time
}

type MenuReport struct {
	MenuID          int
	NamaMenu        string
	Kategori        string
	JumlahTerjual   int
	TotalPendapatan int
}

type Report struct {
	PesananReport []PesananReport
	MenuReport    []MenuReport
}

func (r *Report) GenerateExcel() (*bytes.Buffer, error) {
	f := excelize.NewFile()

	_, err := f.NewSheet("Pesanan")
	if err != nil {
		return nil, err
	}

	_, err = f.NewSheet("Menu")
	if err != nil {
		return nil, err
	}

	pesananHeaders := []string{"Nomor", "Id Pesanan", "Nama Pelanggan", "Tipe Pesanan", "Waktu Pesanan", "Kasir", "Total", "Metode Pembayaran"}
	if err := insertHeaders(f, "Pesanan", 1, pesananHeaders); err != nil {
		return nil, err
	}
	menuHeaders := []string{"Nomor", "Id Menu", "Menu", "Kategori", "Jumlah Terjual", "Total Pendapatan"}
	if err := insertHeaders(f, "Menu", 1, menuHeaders); err != nil {
		return nil, err
	}

	const CELL_A = 65 // Rune for 'A'

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		currCellNum := 2 // Data start on Cell Number 2
		grandTotal := 0
		for i, pesanan := range r.PesananReport {
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A, currCellNum), i+1)
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+1, currCellNum), pesanan.PesananID)
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+2, currCellNum), pesanan.NamaPelanggan)
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+3, currCellNum), pesanan.TipePesanan)
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+4, currCellNum), pesanan.WaktuPesanan)
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+5, currCellNum), pesanan.Kasir)
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+6, currCellNum), pesanan.Total)
			f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+7, currCellNum), pesanan.MetodePembayaran)
			currCellNum++
			grandTotal += pesanan.Total
		}
		f.SetCellValue("Pesanan", fmt.Sprintf("%c%v", CELL_A+6, currCellNum+1), grandTotal)
	}()

	go func() {
		defer wg.Done()
		currCellNum := 2 // Data start on Cell Number 2
		for i, menu := range r.MenuReport {
			f.SetCellValue("Menu", fmt.Sprintf("%c%v", CELL_A, currCellNum), i+1)
			f.SetCellValue("Menu", fmt.Sprintf("%c%v", CELL_A+1, currCellNum), menu.MenuID)
			f.SetCellValue("Menu", fmt.Sprintf("%c%v", CELL_A+2, currCellNum), menu.NamaMenu)
			f.SetCellValue("Menu", fmt.Sprintf("%c%v", CELL_A+3, currCellNum), menu.Kategori)
			f.SetCellValue("Menu", fmt.Sprintf("%c%v", CELL_A+4, currCellNum), menu.JumlahTerjual)
			f.SetCellValue("Menu", fmt.Sprintf("%c%v", CELL_A+5, currCellNum), menu.TotalPendapatan)
			currCellNum++
		}
	}()

	wg.Wait()

	// Delete default Sheet
	f.DeleteSheet("Sheet1")

	// Write to buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return &buf, nil
}

func insertHeaders(f *excelize.File, sheetName string, startRow int, headers []string) error {
	headerStyle, err := f.NewStyle(
		&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
			},
		},
	)

	if err != nil {
		return err
	}

	for i, header := range headers {
		f.SetCellValue(sheetName, fmt.Sprintf("%c%v", 65+i, startRow), header)
	}

	if err := f.SetRowStyle(sheetName, startRow, startRow, headerStyle); err != nil {
		return err
	}

	return nil
}

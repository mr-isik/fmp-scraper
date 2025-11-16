package exporter

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/mr-isik/fmp-scraper/internal/models"
)

// Exporter defines the interface for data export operations
type Exporter interface {
	Export(data []models.StockData, config models.ExportConfig) error
}

// CSVExporter implements CSV export functionality
type CSVExporter struct{}

// NewCSVExporter creates a new CSV exporter
func NewCSVExporter() *CSVExporter {
	return &CSVExporter{}
}

// Export writes stock data to a CSV file
func (e *CSVExporter) Export(data []models.StockData, config models.ExportConfig) error {
	file, err := os.Create(config.OutputFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{
		"Date",
		"Open",
		"High",
		"Low",
		"Close",
		"Adjusted Close",
		"Volume",
		"Unadjusted Volume",
		"Change",
		"Change Percent",
		"VWAP",
		"Label",
		"Change Over Time",
	}

	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write data rows
	for _, item := range data {
		row := []string{
			item.Date,
			formatFloat(item.Open),
			formatFloat(item.High),
			formatFloat(item.Low),
			formatFloat(item.Close),
			formatFloat(item.AdjClose),
			strconv.FormatInt(item.Volume, 10),
			strconv.FormatInt(item.UnadjustedVolume, 10),
			formatFloat(item.Change),
			formatFloat(item.ChangePercent),
			formatFloat(item.VWAP),
			item.Label,
			formatFloat(item.ChangeOverTime),
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}

	return nil
}

// formatFloat converts float64 to string with proper formatting
func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

package models

import "time"

// StockData represents a single stock data point from FMP API
type StockData struct {
	Date             string  `json:"date"`
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Close            float64 `json:"close"`
	AdjClose         float64 `json:"adjClose"`
	Volume           int64   `json:"volume"`
	UnadjustedVolume int64   `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
	VWAP             float64 `json:"vwap"`
	Label            string  `json:"label"`
	ChangeOverTime   float64 `json:"changeOverTime"`
}

// DateRange represents a date range for querying
type DateRange struct {
	From time.Time
	To   time.Time
}

// ExportConfig holds configuration for export operations
type ExportConfig struct {
	OutputFile string
	Symbol     string
	DateRange  DateRange
}

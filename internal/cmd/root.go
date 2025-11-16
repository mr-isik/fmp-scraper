package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mr-isik/fmp-scraper/internal/api"
	"github.com/mr-isik/fmp-scraper/internal/config"
	"github.com/mr-isik/fmp-scraper/internal/exporter"
	"github.com/mr-isik/fmp-scraper/internal/models"
	"github.com/mr-isik/fmp-scraper/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	symbol     string
	fromDate   string
	toDate     string
	outputFile string
	log        = logger.NewConsoleLogger()
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "fmp-scraper",
	Short: "Financial Modeling Prep API data scraper",
	Long: `A powerful CLI tool to fetch financial data from Financial Modeling Prep API 
and export it to CSV format. Supports custom date ranges and output file names.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runScraper(); err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Define flags
	rootCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Stock symbol to fetch (required, e.g., AAPL, TSLA)")
	rootCmd.Flags().StringVarP(&fromDate, "from", "f", "", "Start date in YYYY-MM-DD format (required)")
	rootCmd.Flags().StringVarP(&toDate, "to", "t", "", "End date in YYYY-MM-DD format (required)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output CSV file name (default: {symbol}_{from}_{to}.csv)")

	// Mark required flags
	rootCmd.MarkFlagRequired("symbol")
	rootCmd.MarkFlagRequired("from")
	rootCmd.MarkFlagRequired("to")
}

func runScraper() error {
	// Load configuration
	log.Info("Loading configuration...")
	cfgManager := config.NewEnvConfigManager()
	cfg, err := cfgManager.Load()
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	// Parse dates
	from, err := time.Parse("2006-01-02", fromDate)
	if err != nil {
		return fmt.Errorf("invalid from date format (use YYYY-MM-DD): %w", err)
	}

	to, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		return fmt.Errorf("invalid to date format (use YYYY-MM-DD): %w", err)
	}

	// Validate date range
	if from.After(to) {
		return fmt.Errorf("from date cannot be after to date")
	}

	// Set default output file if not provided
	if outputFile == "" {
		outputFile = fmt.Sprintf("%s_%s_%s.csv", symbol, fromDate, toDate)
	}

	// Create API client
	log.Infof("Fetching data for %s from %s to %s...", symbol, fromDate, toDate)
	apiClient := api.NewFMPClient(cfg.APIKey)

	// Fetch data
	data, err := apiClient.GetHistoricalPrices(symbol, from, to)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}

	log.Infof("Successfully fetched %d records", len(data))

	// Export to CSV
	log.Infof("Exporting data to %s...", outputFile)
	csvExporter := exporter.NewCSVExporter()
	exportConfig := models.ExportConfig{
		OutputFile: outputFile,
		Symbol:     symbol,
		DateRange: models.DateRange{
			From: from,
			To:   to,
		},
	}

	if err := csvExporter.Export(data, exportConfig); err != nil {
		return fmt.Errorf("failed to export data: %w", err)
	}

	log.Success(fmt.Sprintf("Data successfully exported to %s", outputFile))
	return nil
}

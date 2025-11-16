package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/mr-isik/fmp-scraper/internal/models"
)

const (
	baseURL = "https://financialmodelingprep.com/api/v3"
	timeout = 30 * time.Second
)

// Client defines the interface for FMP API operations
type Client interface {
	GetHistoricalPrices(symbol string, from, to time.Time) ([]models.StockData, error)
	GetQuote(symbol string) ([]models.StockData, error)
}

// HTTPClient defines the interface for HTTP operations
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// FMPClient implements the Client interface
type FMPClient struct {
	apiKey     string
	httpClient HTTPClient
}

// NewFMPClient creates a new FMP API client
func NewFMPClient(apiKey string) *FMPClient {
	return &FMPClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// NewFMPClientWithHTTP creates a new FMP API client with custom HTTP client (for testing)
func NewFMPClientWithHTTP(apiKey string, httpClient HTTPClient) *FMPClient {
	return &FMPClient{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// GetHistoricalPrices fetches historical stock prices for a given symbol and date range
func (c *FMPClient) GetHistoricalPrices(symbol string, from, to time.Time) ([]models.StockData, error) {
	endpoint := fmt.Sprintf("%s/historical-price-full/%s", baseURL, symbol)

	params := url.Values{}
	params.Add("apikey", c.apiKey)
	params.Add("from", from.Format("2006-01-02"))
	params.Add("to", to.Format("2006-01-02"))

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Symbol     string             `json:"symbol"`
		Historical []models.StockData `json:"historical"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result.Historical) == 0 {
		return nil, fmt.Errorf("no data found for symbol %s", symbol)
	}

	return result.Historical, nil
}

// GetQuote fetches the current quote for a given symbol
func (c *FMPClient) GetQuote(symbol string) ([]models.StockData, error) {
	endpoint := fmt.Sprintf("%s/quote/%s", baseURL, symbol)

	params := url.Values{}
	params.Add("apikey", c.apiKey)

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result []models.StockData
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no data found for symbol %s", symbol)
	}

	return result, nil
}

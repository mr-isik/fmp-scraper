package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	APIKey string
}

// Manager defines the interface for configuration management
type Manager interface {
	Load() (*Config, error)
	Validate() error
}

// EnvConfigManager implements configuration loading from environment
type EnvConfigManager struct {
	config *Config
}

// NewEnvConfigManager creates a new environment config manager
func NewEnvConfigManager() *EnvConfigManager {
	return &EnvConfigManager{}
}

// Load loads configuration from environment variables
func (m *EnvConfigManager) Load() (*Config, error) {
	// Try to load .env file (optional)
	_ = godotenv.Load()

	apiKey := os.Getenv("FMP_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("FMP_API_KEY environment variable is required")
	}

	m.config = &Config{
		APIKey: apiKey,
	}

	return m.config, nil
}

// Validate validates the configuration
func (m *EnvConfigManager) Validate() error {
	if m.config == nil {
		return fmt.Errorf("configuration not loaded")
	}

	if m.config.APIKey == "" {
		return fmt.Errorf("API key is required")
	}

	return nil
}

// GetConfig returns the loaded configuration
func (m *EnvConfigManager) GetConfig() *Config {
	return m.config
}

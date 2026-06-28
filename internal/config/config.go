package config

import (
	"fmt"
	"os"
)

type Config struct {
	GoogleCredentials string
	SpreadsheetID     string
	SheetName         string

	APIURL   string
	APIToken string

	UserName string
}

func Load() (*Config, error) {

	cfg := &Config{
		SpreadsheetID:     os.Getenv("SPREADSHEET_ID"),
		SheetName:         os.Getenv("SHEET_NAME"),

		APIURL:   os.Getenv("API_URL"),
		APIToken: os.Getenv("API_TOKEN"),

		UserName: os.Getenv("USER_NAME"),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.SpreadsheetID == "" {
		return fmt.Errorf("SPREADSHEET_ID not found")
	}

	if c.APIURL == "" {
		return fmt.Errorf("API_URL not found")
	}

	if c.UserName == "" {
		return fmt.Errorf("USER_NAME not found")
	}

	return nil
}
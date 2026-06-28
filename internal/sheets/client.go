package sheets

import (
	"context"
	"fmt"
	"time"

	"github.com/yssacst/google-sheets-bot/internal/config"
	"github.com/yssacst/google-sheets-bot/internal/core"

	"google.golang.org/api/sheets/v4"
)

type Client struct {
	service       *sheets.Service
	spreadsheetID string
	sheetName     string
}

func NewClient(cfg *config.Config) (*Client, error) {

	ctx := context.Background()

	service, err := sheets.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create sheets service: %w", err)
	}

	return &Client{
		service:       service,
		spreadsheetID: cfg.SpreadsheetID,
		sheetName:     cfg.SheetName,
	}, nil
}

func (c *Client) GetRows(ctx context.Context) ([]core.Row, error) {

	var rows []core.Row

	resp, err := c.service.Spreadsheets.Values.Get(
		c.spreadsheetID,
		c.sheetName,
	).Context(ctx).Do()

	if err != nil {
		return nil, fmt.Errorf("failed to read sheet: %w", err)
	}

	if len(resp.Values) == 0 {
		return []core.Row{}, nil
	}

	// assumindo formato:
	// Nome | Data
	for i, r := range resp.Values {

		// pula header
		if i == 0 {
			continue
		}

		row := core.Row{}

		if len(r) > 0 {
			row.Name = fmt.Sprintf("%v", r[0])
		}

		if len(r) > 1 {
			parsedDate, _ := parseSheetDate(r[1])
			row.Date = parsedDate
		}

		rows = append(rows, core.Row{
			Name: row.Name,
			Date: row.Date,
		})
	}

	return rows, nil
}

func parseSheetDate(v any) (time.Time, error) {

	switch value := v.(type) {

	case string:
		// ISO
		t, err := time.Parse("2006-01-02", value)
		if err == nil {
			return t, nil
		}
		// BR
		return time.Parse("02/01/2006", value)
	case float64:
		// Google Sheets serial date
		return timeFromExcelSerial(value), nil
	default:
		return time.Time{}, fmt.Errorf("unsupported date format: %v", v)
	}
}

func timeFromExcelSerial(serial float64) time.Time {
	// Google Sheets base date
	base := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	return base.AddDate(0, 0, int(serial))
}
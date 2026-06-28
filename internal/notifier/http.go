package notifier

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/yssacst/google-sheets-bot/internal/config"
)

type Client struct {
	apiURL   string
	apiToken string
	http     *http.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		apiURL:   cfg.APIURL,
		apiToken: cfg.APIToken,
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Send(ctx context.Context, payload Payload) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.apiURL,
		strings.NewReader(payload.Message),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Markdown", "true")

	if payload.Title != "" {
		req.Header.Set("Title", payload.Title)
	}

	if payload.Priority > 0 {
		req.Header.Set("Priority", fmt.Sprint(payload.Priority))
	}

	if len(payload.Tags) > 0 {
		req.Header.Set("Tags", strings.Join(payload.Tags, ","))
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
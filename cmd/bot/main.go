package main

import (
	"context"
	"fmt"
	"log"
	"time"
	
	"github.com/yssacst/google-sheets-bot/internal/config"
	"github.com/yssacst/google-sheets-bot/internal/core"
	"github.com/yssacst/google-sheets-bot/internal/logger"
	"github.com/yssacst/google-sheets-bot/internal/notifier"
	"github.com/yssacst/google-sheets-bot/internal/sheets"
	"github.com/yssacst/google-sheets-bot/internal/version"
	
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	lg := logger.New()

	ctx := context.Background()

	version.GetInfo()

	// Load config
	cfg, err := config.Load()
	if err != nil {
		lg.Error(fmt.Sprintf("failed to load config: %v", err))
		log.Fatal(err)
	}

	// Init adapters
	sheetClient, err := sheets.NewClient(cfg)
	if err != nil {
		lg.Error(fmt.Sprintf("failed to init sheets client: %v", err))
		log.Fatal(err)
	}

	apiClient := notifier.NewClient(cfg)

	// Run flow
	err = run(ctx, cfg, sheetClient, apiClient, lg)
	if err != nil {
		lg.Error(fmt.Sprintf("bot execution failed: %v", err))
		log.Fatal(err)
	}

	lg.Info("bot finished successfully")
}

func run(
	ctx context.Context,
	cfg *config.Config,
	sheetClient *sheets.Client,
	apiClient *notifier.Client,
	lg *logger.Logger,
) error {

	rows, err := sheetClient.GetRows(ctx)
	if err != nil {
		return err
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	today := time.Now().In(loc)

	isOnDuty := core.IsOnDutyTomorrow(rows, cfg.UserName, today)

	if !isOnDuty {
		lg.Info(fmt.Sprintf("user %v is NOT on duty tomorrow", cfg.UserName))
		return nil
	}

	lg.Info(fmt.Sprintf("user %v IS on duty tomorrow → sending notification", cfg.UserName))

	payload := notifier.BuildPayload(cfg.UserName)

	err = apiClient.Send(ctx, payload)

	if err != nil {
		return err
	}

	return nil
}
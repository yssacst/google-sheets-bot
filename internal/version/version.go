package version

import "github.com/yssacst/google-sheets-bot/internal/logger"

var (
	Version = "dev"
	Commit  = "unknown"
)

func GetInfo()  {
	lg := logger.New()
	lg.Info(
		"Version:",
		Version,
		"; Commit: ",
		Commit,
	)
}
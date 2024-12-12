package logging

import (
	"log/slog"
	"os"
	"path"
)

func SetupLogger() (*os.File, error) {
	logsDir, err := os.MkdirTemp(os.TempDir(), "tuner-logs-*")
	if err != nil {
		return nil, err
	}

	logFilePath := path.Join(logsDir, "log.txt")
	logFile, err := os.Create(logFilePath)
	if err != nil {
		return nil, err
	}

	handler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("oh, hello there", "logFilePath", logFilePath)

	return logFile, nil
}

package logging

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"
)

func SetupLogger() (*os.File, error) {
	logsDir, err := os.MkdirTemp(os.TempDir(), fmt.Sprintf("tuner-logs-%d", time.Now().UnixMilli()))
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

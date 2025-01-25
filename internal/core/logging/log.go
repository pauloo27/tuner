package logging

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"
)

func SetupLogger() (*os.File, error) {
	logDir := path.Join(os.TempDir(), "tuner-logs")
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	logFilePath := path.Join(logDir, fmt.Sprintf("%d-%d.txt", time.Now().Unix(), os.Getpid()))
	/* #nosec G304 */
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

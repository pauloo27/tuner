package logging

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"

	slogmulti "github.com/samber/slog-multi"
)

var MemoryBuffer bytes.Buffer

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

	fileHandler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	memoryHandler := slog.NewTextHandler(&MemoryBuffer, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	combinedHandler := slogmulti.Fanout(fileHandler, memoryHandler)

	logger := slog.New(combinedHandler)
	slog.SetDefault(logger)

	slog.Info("oh, hello there", "logFilePath", logFilePath)

	return logFile, nil
}

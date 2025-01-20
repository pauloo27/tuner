package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/pauloo27/tuner/internal/core/logging"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/ui"

	"github.com/pauloo27/tuner/internal/providers/player/mpv"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/providers/source/yt"
)

func main() {
	fmt.Println("Hello world!")
	logFile, err := logging.SetupLogger()
	if err != nil {
		slog.Info("Failed to setup logger!", "err", err)
		// hopefully, the only panic in the whole code base! (it's not)
		panic(err)
	}
	defer onAppClose(logFile)

	initProviders()

	if err := ui.StartTUI(); err != nil {
		slog.Error("Failed to start TUI", "err", err)
		os.Exit(2)
	}
}

func initProviders() {
	mpvPlayer, err := mpv.NewMpvPlayer()
	if err != nil {
		slog.Error("Failed to init mpv player!", "err", err)
		os.Exit(1)
	}

	providers.Player = mpvPlayer
	providers.Sources = []source.Source{yt.NewYoutubeSource()}

	sourcesNames := make([]string, 0, len(providers.Sources))

	for _, source := range providers.Sources {
		sourcesNames = append(sourcesNames, source.Name())
	}

	slog.Info("Player provider", "name", providers.Player.Name())
	slog.Info("Sources", "names", sourcesNames)
}

func onAppClose(logFile *os.File) {
	if err := recover(); err != nil {
		// FIXME: the stacktrace is unreadable this way
		slog.Error("PANIC!", "err", err, "stacktrace", debug.Stack())
		fmt.Println("Panic caught! Loggin' and exiting...")
	} else {
		slog.Info("Goodbye cruel world!")
		fmt.Println("Goodbye cruel world!")
	}

	err := logFile.Close()
	if err == nil {
		fmt.Printf("Log saved to %s\n", logFile.Name())
	} else {
		fmt.Printf("Failed to close log file: %v", err)
	}
}

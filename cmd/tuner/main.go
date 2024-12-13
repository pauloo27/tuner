package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/pauloo27/tuner/internal/core/logging"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/pages/home"
	"github.com/pauloo27/tuner/internal/ui/pages/playing"
	"github.com/pauloo27/tuner/internal/ui/pages/searching"

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
	if err = registerPages(); err != nil {
		slog.Error("Failed to register page", "err", err)
		os.Exit(-1)
	}

	if err = ui.StartTUI(); err != nil {
		slog.Error("Failed to run TUI", "err", err)
		os.Exit(-1)
	}
}

func registerPages() error {
	ui.Setup()
	return ui.RegisterPages(
		home.NewHomePage(), searching.NewSearchingPage(), playing.NewPlayingPage(),
	)
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
		sourcesNames = append(sourcesNames, source.GetName())
	}

	slog.Info("Player provider", "name", providers.Player.GetName())
	slog.Info("Sources", "names", sourcesNames)
}

func onAppClose(logFile *os.File) {
	if err := recover(); err != nil {
		slog.Error("PANIC!", "err", err, "stacktrace", debug.Stack())
		fmt.Println("Panic caught! Loggin and exiting...")
	} else {
		slog.Info("Goodbye cruel world!")
		fmt.Println("Goodbye cruel world!")
	}

	_ = logFile.Close()
	fmt.Printf("Log saved to %s\n", logFile.Name())
}

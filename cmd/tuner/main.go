package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pauloo27/tuner/internal/core/logging"
	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/ui"

	_ "github.com/pauloo27/tuner/internal/ui/pages/home"
	_ "github.com/pauloo27/tuner/internal/ui/pages/playing"
	_ "github.com/pauloo27/tuner/internal/ui/pages/searching"

	"github.com/pauloo27/tuner/internal/providers/player/mpv"
	"github.com/pauloo27/tuner/internal/providers/source"
	"github.com/pauloo27/tuner/internal/providers/source/yt"
)

func main() {
	logFile, err := logging.SetupLogger()
	if err != nil {
		slog.Info("Failed to setup logger!", "err", err)
		// hopefully, the only panic in the whole code base! (it's not)
		panic(err)
	}
	defer onAppClose(logFile)

	initProviders()

	err = ui.StartApp("home")
	if err != nil {
		slog.Error("Failed to start app", "err", err)
		os.Exit(1)
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
		sourcesNames = append(sourcesNames, source.GetName())
	}

	slog.Info("Player provider", "name", providers.Player.GetName())
	slog.Info("Sources", "names", sourcesNames)
}

func onAppClose(logFile *os.File) {
	_ = logFile.Close()
	slog.Info("Goodbye!")
	fmt.Println("Goodbye!")
	fmt.Printf("Log saved to %s\n", logFile.Name())
}

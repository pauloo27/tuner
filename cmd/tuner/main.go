package main

import (
	"log/slog"
	"os"

	"github.com/pauloo27/tuner/internal/core/logging"
	"github.com/pauloo27/tuner/internal/ui"

	_ "github.com/pauloo27/tuner/internal/ui/pages/home"
	_ "github.com/pauloo27/tuner/internal/ui/pages/playing"
	_ "github.com/pauloo27/tuner/internal/ui/pages/searching"

	"github.com/pauloo27/tuner/internal/providers/player/mpv"
)

func main() {
	logFile, err := logging.SetupLogger()
	if err != nil {
		slog.Info("Failed to setup logger!", "err", err)
		// hopefully, the only panic in the whole code base! (it's not)
		panic(err)
	}
	defer logFile.Close()

	err = mpv.InitMpvPlayer()
	if err != nil {
		slog.Error("Failed to init mpv player!", "err", err)
		os.Exit(-1)
	}

	err = ui.StartApp("home")
	if err != nil {
		slog.Error("Failed to start app", "err", err)
		os.Exit(1)
	}
}

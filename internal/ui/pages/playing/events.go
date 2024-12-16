package playing

import (
	"log/slog"

	"github.com/pauloo27/tuner/internal/providers"
	"github.com/pauloo27/tuner/internal/providers/player"
	"github.com/pauloo27/tuner/internal/ui"
	"github.com/pauloo27/tuner/internal/ui/core"
)

func (p *playingPage) registerListeners() {
	// TODO: should I ignore the events when the page is not focused?
	providers.Player.On(player.PlayerEventPause, func(...any) {
		ui.App.QueueUpdateDraw(func() {
			p.songLabel.SetText(p.buildSongLabel(core.IconPaused))
		})
	})

	providers.Player.On(player.PlayerEventPlay, func(...any) {
		ui.App.QueueUpdateDraw(func() {
			p.songLabel.SetText(p.buildSongLabel(core.IconPlaying))
		})
	})

	providers.Player.On(player.PlayerEventVolumeChanged, func(...any) {
		ui.App.QueueUpdateDraw(func() {
			err := p.updateVolumeLabel()
			if err != nil {
				slog.Info("Failed to update volume", "err", err)
			}
		})
	})

	providers.Player.On(player.PlayerEventIdle, func(...any) {
		ui.App.QueueUpdateDraw(func() {
			ui.SwitchPage(ui.HomePageName)
		})
	})
}

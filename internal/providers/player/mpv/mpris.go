package mpv

func (p *MpvPlayer) loadMpris() error {
	// TODO: better
	path := "/etc/mpv/scripts/mpris.so"

	return p.instance.Command([]string{"load-script", path})
}

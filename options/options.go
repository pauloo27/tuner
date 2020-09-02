package options

type TunerOptions struct {
	KeepLiveCache, ShowVideo bool
}

var Options = TunerOptions{
	ShowVideo:     false,
	KeepLiveCache: false,
}

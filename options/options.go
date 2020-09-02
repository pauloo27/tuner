package options

type TunerOptions struct {
	Cache, ShowVideo bool
}

var Options = TunerOptions{
	ShowVideo: false,
	Cache:     false,
}

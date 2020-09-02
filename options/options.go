package options

type TunerOptions struct {
	KeepCacheFromLives, ShowVideo bool
}

var Options = TunerOptions{
	ShowVideo:          false,
	KeepCacheFromLives: false,
}

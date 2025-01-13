package progress

type ProgressBarStyle interface {
	Draw(bar *ProgressBar)
}

package mode

type Mode struct {
	Displayed bool
	Handler   func()
}

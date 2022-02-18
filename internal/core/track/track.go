package track

type Track interface {
	Title() string
	Artist() string
	GetMediaURL() string
}

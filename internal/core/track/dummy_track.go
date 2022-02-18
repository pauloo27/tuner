package track

type DummyTrack struct {
	TrackTitle, TrackArtist, MediaURL string
}

var _ Track = DummyTrack{}

func (d DummyTrack) Artist() string {
	return d.TrackArtist
}

func (d DummyTrack) Title() string {
	return d.TrackTitle
}

func (d DummyTrack) GetMediaURL() string {
	return d.MediaURL
}

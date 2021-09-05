package search

type YouTubeSearch struct {
}

var _ SearchProvider = &YouTubeSearch{}

func (YouTubeSearch) GetName() (name string) {
	return "YouTube"
}

func (YouTubeSearch) SearchFor(searchQuery string) []*SearchResult {
	// TODO:
	return nil
}

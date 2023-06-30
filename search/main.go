package search

type SearchResult struct {
	Title, Uploader, URL, Duration, ID string
	Live                               bool
	SourceName                         string
	Extra                              []string
}

type SearchSource interface {
	Search(query string, limit int) []*SearchResult
}

var (
	SourceYouTube = YouTubeSource{}
)

func Search(query string, limit int, sources ...SearchSource) []*SearchResult {
	results := []*SearchResult{}
	sourcesCount := len(sources)
	if sourcesCount == 0 {
		return results
	}
	for _, source := range sources {
		results = append(results, source.Search(query, limit/sourcesCount)...)
	}
	return results
}

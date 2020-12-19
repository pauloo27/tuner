package search

type SearchResult struct {
	Title, Uploader, URL, Duration, ID string
	Live                               bool
	SourceName                         string
}

type SearchSource interface {
	Search(query string, limit int) []*SearchResult
}

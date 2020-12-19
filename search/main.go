package search

type SearchResult struct {
	Title, Uploader, URL, Duration string
	Live                           bool
	SourceName                     string
}

type SearchSource interface {
	Search(query string, limit int) []*SearchResult
}

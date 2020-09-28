package structures

// Request : to unmarshal Incoming request data
type Request struct {
	URL              string `json:"url"`
	CrawlTimeout     int    `json:"crawl_timeout"`
	Frequency        int    `json:"frequency"`
	FailureThreshold int    `json:"failure_threshold"`
}

package concurrency

// WebsiteChecker checks a url, returning a bool
type WebsiteChecker func(string) bool

// URLCheckResult ...
type URLCheckResult struct {
	URL    string
	Result bool
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map
// of urls to the result of checking each url with the WebsiteChecker function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	checksChannel := make(chan URLCheckResult)

	for _, url := range urls {
		go func(u string) {
			checksChannel <- URLCheckResult{
				URL:    u,
				Result: wc(u),
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		check := <-checksChannel
		results[check.URL] = check.Result
	}

	return results
}

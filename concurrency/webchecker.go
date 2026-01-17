package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, websites []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result)

	// without go routine
	// for _, website := range websites {
	// 	results[website] = wc(website)
	// }

	// with go routine
	for _, website := range websites {
		go func() {
			resultChannel <- result{website, wc(website)}
		}()
	}

	for range websites {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

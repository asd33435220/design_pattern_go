package main

type nginx struct {
	appliacation      *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		appliacation:      &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Now Allowed"
	}
	return n.appliacation.handleRequest(url, method)
}

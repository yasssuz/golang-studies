package main

import (
	"fmt"
	"net/http"
)

type fetcherResult struct {
	url      string
	error    error
	response *http.Response
}

func main() {
	// Error handling with concurrent code gets tricky pretty quickly: who's responsible for handling it?
	// Because concurrent processes are operating independently of its parents or siblings, those same concurrent
	// processes should tightly couple the error to the result type, allowing other parts of the program to make more
	// informed decisions about the errors.

	done := make(chan any)
	fetcherStreamOwner := func(done <-chan any, urls []string) <-chan fetcherResult {
		fetches := make(chan fetcherResult)
		go func() {
			defer close(fetches)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := fetcherResult{
					url:      url,
					error:    err,
					response: resp,
				}
				select {
				case <-done:
					return
				case fetches <- result:
				}
			}
		}()
		return fetches
	}
	fetcherStreamConsumer := func(fetches <-chan fetcherResult) {
		for fetch := range fetches {
			if fetch.error != nil {
				fmt.Println(fmt.Errorf("Failed to fetch %s with error %w ", fetch.url, fetch.error))
				return
			}
			fmt.Println(fetch)
		}
	}

	urls := []string{"https://www.google.com", "https://badhost"}
	fetches := fetcherStreamOwner(done, urls)
	fetcherStreamConsumer(fetches)
}

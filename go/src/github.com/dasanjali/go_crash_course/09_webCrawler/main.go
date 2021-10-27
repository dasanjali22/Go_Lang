/*In this exercise you'll use Go's concurrency features to parallelize a web crawler.
Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use! */

package main

import (
	"fmt"
)

type Fetcher interface {
	//  The fetch  will get return the body of URL
	// and a cut of URL's found on that page.
	Fetch(url_link string) (body string, urls []string, err error)
}

type FetchedUrl struct {
	parent_url string
	body       string
	urls       []string
	err_state  error
	depth      int
}

type fake_Url_Fetcher map[string]*fakeResult_Resp

type fakeResult_Resp struct {
	body string
	urls []string
}

// Crawl function uses the fetcher to recursively do the crawling
// for pages starting with url, to a maximum of depth.
func Crawl(url_link string, depth int, fetcher Fetcher) {
	results := make(chan *FetchedUrl)
	alreadyFetched_Url := make(map[string]bool)

	fetch := func(url_link string, depth int) {
		body, urls, err_state := fetcher.Fetch(url_link)
		results <- &FetchedUrl{url_link, body, urls, err_state, depth}
	}

	go fetch(url_link, depth)
	alreadyFetched_Url[url_link] = true

	for countFetching := 1; countFetching > 0; countFetching-- {
		result := <-results

		if result.err_state != nil { // it will skip wrong urls
			fmt.Println(result.err_state)
			continue
		}

		fmt.Printf("url found: %s %q\n", result.parent_url, result.body)

		if result.depth > 0 {
			for _, newUrl := range result.urls {
				if !alreadyFetched_Url[newUrl] {
					countFetching++
					go fetch(newUrl, depth-1)
					alreadyFetched_Url[newUrl] = true
				}
			}
		}
	}

	close(results)
}

func main() {
	Crawl("http://golang.org/", 7, fetcher)
}

func (fake fake_Url_Fetcher) Fetch(web_url string) (string, []string, error) {
	if res, ok := fake[web_url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("url not found: %s", web_url)
}

var fetcher = fake_Url_Fetcher{
	"http://golang.org/": &fakeResult_Resp{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult_Resp{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult_Resp{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult_Resp{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

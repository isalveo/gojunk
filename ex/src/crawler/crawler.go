package crawler

import (
  "fmt"
)

type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}

var store map[string]bool = make(map[string]bool)

func Krawl(url string, fetcher Fetcher, Urls chan []string) {
  body, urls, err := fetcher.Fetch(url)

  if err != nil {
    fmt.Printf("not found: %s", url)
  } else {
    fmt.Printf("found %s: %q\n", url, body)
  }

  Urls <- urls
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  Urls := make(chan []string)
  go Krawl(url, fetcher, Urls)

  band := 1
  store[url] = true

  for i := 0; i < depth; i++ {

    for band > 0 {
      band--
      next := <- Urls

      for _, u := range next {
        if _, done := store[u]; !done {
          store[u] = true
          band++
          go Krawl(u, fetcher, Urls)
        }
      }
    }
  }
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
  body string
  urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
  if res, ok := f[url]; ok {
    return res.body, res.urls, nil
  }
  return "", nil, fmt.Errorf("not found: %s", url)
}

var FetcherT = fetcher

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
  "http://golang.org/": &fakeResult{
    "The Go Programming Language",
    []string{
      "http://golang.org/pkg/",
      "http://golang.org/cmd/",
    },
  },
  "http://golang.org/pkg/": &fakeResult{
    "Packages",
    []string{
      "http://golang.org/",
      "http://golang.org/cmd/",
      "http://golang.org/pkg/fmt/",
      "http://golang.org/pkg/os/",
    },
  },
  "http://golang.org/pkg/fmt/": &fakeResult{
    "Package fmt",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
  "http://golang.org/pkg/os/": &fakeResult{
    "Package os",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
}

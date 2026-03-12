package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://ardanlabs.com",
		"https://ibm.com/no/such/page",
	}

	start := time.Now()
	/*
		for _, url := range urls {
			status, err := urlCheck(url)
			fmt.Printf("%q: %d (%v)\n", url, status, err)
		}
	*/
	//fanOutResult(urls)
	fanOutWait(urls)
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
}

func urlLog(url string) {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("urlLog", "url", url, "error", err)
		return
	}

	slog.Info("urlLog", "url", url, "status", resp.StatusCode)
}

func fanOutWait(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			urlLog(url)
		}()
	}

	wg.Wait()
}

func fanOutResult(urls []string) {
	type result struct {
		url    string
		status int
		err    error
	}

	ch := make(chan result)
	for _, url := range urls {
		go func() {
			r := result{url: url}
			defer func() { ch <- r }()

			r.status, r.err = urlCheck(url)
		}()
	}

	for range urls {
		r := <-ch
		fmt.Printf("%q: %d (%v)\n", r.url, r.status, r.err)
	}

}

func urlCheck(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

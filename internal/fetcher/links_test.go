package fetcher

import (
	"fmt"
	"testing"
)

func TestGetLinks(t *testing.T) {
	links, err := getLinksRss("https://newtalk.tw/rss/all/")
	if err != nil {
		t.Error(err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}

func TestFetchLinks(t *testing.T) {
	links, err := fetchLinks()
	if err != nil {
		t.Error(err)
	}

	for _, link := range links {
		fmt.Println(link)
	}
	fmt.Println(len(links))
}

package fetcher

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/hi20160616/exhtml"
	"github.com/pkg/errors"
)

func TestFetchTitle(t *testing.T) {
	tests := []struct {
		url   string
		title string
	}{
		{"https://newtalk.tw/news/view/2021-11-02/660037", "批藍反萊豬鎖死台灣 蘇貞昌：朱立倫講瘋豬肉還想親美？ | 政治"},
	}
	for _, tc := range tests {
		a := NewArticle()
		u, err := url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		a.U = u
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		got, err := a.fetchTitle()
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			} else {
				fmt.Println("ignore pass test: ", tc.url)
			}
		} else {
			if tc.title != got {
				t.Errorf("\nwant: %s\n got: %s", tc.title, got)
			}
		}
	}

}

func TestFetchUpdateTime(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"https://newtalk.tw/news/view/2021-11-02/660037", "2021-11-02 09:59:07 +0800 UTC"},
	}
	var err error
	for _, tc := range tests {
		a := NewArticle()
		a.U, err = url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		tt, err := a.fetchUpdateTime()
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			}
		}
		ttt := tt.AsTime()
		got := shanghai(ttt)
		if got.String() != tc.want {
			t.Errorf("\nwant: %s\n got: %s", tc.want, got.String())
		}
	}
}

func TestFetchContent(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"https://newtalk.tw/news/view/2021-11-02/660037", "2021-11-02 09:59:07 +0800 UTC"},
	}
	var err error

	for _, tc := range tests {
		a := NewArticle()
		a.U, err = url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		c, err := a.fetchContent()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(c)
	}
}

func TestFetchArticle(t *testing.T) {
	tests := []struct {
		url string
		err error
	}{
		{"https://newtalk.tw/news/view/2021-10-28/658199", ErrTimeOverDays},
		{"https://newtalk.tw/news/view/2021-11-02/660037", nil},
	}
	for _, tc := range tests {
		a := NewArticle()
		a, err := a.fetchArticle(tc.url)
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			} else {
				fmt.Println("ignore old news pass test: ", tc.url)
			}
		} else {
			fmt.Println("pass test: ", a.Content)
		}
	}
}

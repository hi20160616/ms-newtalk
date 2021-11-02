package service

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/hi20160616/fetchnews-api/proto/v1"
)

func TestListArticles(t *testing.T) {
	s := &Server{}
	ss, err := s.ListArticles(context.Background(), &v1.ListArticlesRequest{})
	if err != nil {
		t.Error(err)
	}
	for _, e := range ss.Articles {
		fmt.Println(e)
	}
}

func TestGetArticle(t *testing.T) {
	s := &Server{}
	ss, err := s.GetArticle(
		context.Background(),
		&v1.GetArticleRequest{Id: "03418273947128734"})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ss.Title)
}

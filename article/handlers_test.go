package article

import (
	"net/http"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)

	testPageUnauthenticated(t, req, "<title>Home Page</title>", "GET", "/", showIndexPage)
}

func TestArticleUnauthenticated(t *testing.T) {

	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testPageUnauthenticated(t, req, "<title>Article 1</title>", "GET", "/article/view/:id", getArticle)
}

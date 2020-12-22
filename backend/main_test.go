package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/wiki/backend/controllers"
)

type routesFunc func(res http.ResponseWriter, req *http.Request)

func TestNotAllowedMethod(t *testing.T) {
	var server routesFunc
	server = controllers.NewWikiHandlers().ArticleRoutes

	t.Run("returns 405 NotAllowedMethod", func(t *testing.T) {
		reqPost, _ := http.NewRequest(http.MethodPost, "/articles/", nil)
		resPost := httptest.NewRecorder()

		server(resPost, reqPost)
		postStatusCode := resPost.Result().StatusCode
		assertResponseStatuscode(t, postStatusCode, http.StatusMethodNotAllowed)

		reqDelete, _ := http.NewRequest(http.MethodPost, "/articles/", nil)
		resDelete := httptest.NewRecorder()

		server(resDelete, reqDelete)
		deleteStatusCode := resDelete.Result().StatusCode
		assertResponseStatuscode(t, deleteStatusCode, http.StatusMethodNotAllowed)
	})
}

func TestListStoredArticlesAreEmpty(t *testing.T) {
	var server routesFunc
	server = controllers.NewWikiHandlers().ArticleRoutes

	t.Run("it returns empty array", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/articles/", nil)
		res := httptest.NewRecorder()

		server(res, req)

		statusCode := res.Result().StatusCode
		assertResponseStatuscode(t, statusCode, http.StatusOK)

		contentType := res.Result().Header.Get("content-type")
		if contentType != "application/json" {
			t.Errorf("expected contentType application/json; but got %v", contentType)
		}

		var got []string
		json.NewDecoder(res.Body).Decode(got)

		if len(got) > 0 {
			t.Errorf("expected an empty array but got an array that has elements")
		}
	})
}

func TestSingleArticleIsEmpty(t *testing.T) {
	t.Run("returns 404 Not Found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/articles/wiki", nil)
		res := httptest.NewRecorder()

		controllers.NewWikiHandlers().ArticleRoutes(res, req)

		statusCode := res.Result().StatusCode
		assertResponseStatuscode(t, statusCode, http.StatusNotFound)
	})
}

func TestMutateArticle(t *testing.T) {
	var routes routesFunc
	routes = controllers.NewWikiHandlers().ArticleRoutes

	newArticle := "A wiki is a knowledge base website"
	updatedArticle := "A wiki is a best website"

	t.Run("it returns 201 Created when an aricle is created and it returns 200 OK when the article is updated", func(t *testing.T) {
		// Create an article
		ioReader := strings.NewReader(newArticle)
		reqCreateArticle, _ := http.NewRequest(http.MethodPut, "/articles/wiki", ioReader)
		resCreateArticle := httptest.NewRecorder()
		routes(resCreateArticle, reqCreateArticle)
		statusCode := resCreateArticle.Result().StatusCode

		assertResponseStatuscode(t, statusCode, http.StatusCreated)

		got := resCreateArticle.Body.String()

		if len(got) != 0 {
			t.Errorf("expected no payload; but got %v", got)
		}

		// Update the article
		ioReaderUpdate := strings.NewReader(updatedArticle)
		reqUpdateArticle, _ := http.NewRequest(http.MethodPut, "/articles/wiki", ioReaderUpdate)
		resUpdateArticle := httptest.NewRecorder()

		routes(resUpdateArticle, reqUpdateArticle)

		statusCodeUpdateArticle := resUpdateArticle.Result().StatusCode
		assertResponseStatuscode(t, statusCodeUpdateArticle, http.StatusOK)
	})
}

func newGetArticleRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/articles/%s", name), nil)
	return req
}

func assertResponseStatuscode(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("expected status %v; but got %v", expected, got)
	}
}

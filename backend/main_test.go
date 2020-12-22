package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/wiki/backend/controllers"
)

type routesFunc func(res http.ResponseWriter, req *http.Request)

func TestNotAllowedMethod(t *testing.T) {
	var serveHTTP routesFunc
	serveHTTP = controllers.NewWikiHandlers().ArticleRoutes

	t.Run("returns 405 NotAllowedMethod", func(t *testing.T) {
		reqPost, _ := http.NewRequest(http.MethodPost, "/articles/", nil)
		resPost := httptest.NewRecorder()

		serveHTTP(resPost, reqPost)
		postStatusCode := resPost.Result().StatusCode
		assertResponseStatuscode(t, postStatusCode, http.StatusMethodNotAllowed)

		reqDelete, _ := http.NewRequest(http.MethodPost, "/articles/", nil)
		resDelete := httptest.NewRecorder()

		serveHTTP(resDelete, reqDelete)
		deleteStatusCode := resDelete.Result().StatusCode
		assertResponseStatuscode(t, deleteStatusCode, http.StatusMethodNotAllowed)
	})
}

func TestListStoredArticlesAreEmpty(t *testing.T) {
	var serveHTTP routesFunc
	serveHTTP = controllers.NewWikiHandlers().ArticleRoutes

	t.Run("it returns empty array", func(t *testing.T) {
		req := newGetArticleRequest("")
		res := httptest.NewRecorder()

		serveHTTP(res, req)

		statusCode := res.Result().StatusCode
		assertResponseStatuscode(t, statusCode, http.StatusOK)

		gotContentType := res.Result().Header.Get("content-type")
		assertContentType(t, gotContentType, "application/json")

		var got []string
		json.NewDecoder(res.Body).Decode(got)

		if len(got) > 0 {
			t.Errorf("expected an empty array but got an array that has elements")
		}
	})
}

func TestStoreArticle(t *testing.T) {
	var routes routesFunc
	routes = controllers.NewWikiHandlers().ArticleRoutes

	newArticle := "A wiki is a knowledge base website"
	updatedArticle := "A wiki is a best website"

	t.Run("it returns 201 Created when an aricle is created", func(t *testing.T) {
		// Create an article
		ioReader := strings.NewReader(newArticle)
		reqCreateArticle := newPutArticleRequest("wiki", ioReader)
		resCreateArticle := httptest.NewRecorder()
		routes(resCreateArticle, reqCreateArticle)
		statusCode := resCreateArticle.Result().StatusCode

		assertResponseStatuscode(t, statusCode, http.StatusCreated)

		got := resCreateArticle.Body.String()

		assertResponsePayload(t, got, "expected no payload; but got %v")
	})

	t.Run("it returns 200 OK when the article is updated", func(t *testing.T) {
		// Update the article
		ioReaderUpdate := strings.NewReader(updatedArticle)
		reqUpdateArticle := newPutArticleRequest("wiki", ioReaderUpdate)
		resUpdateArticle := httptest.NewRecorder()

		routes(resUpdateArticle, reqUpdateArticle)

		statusCodeUpdateArticle := resUpdateArticle.Result().StatusCode
		assertResponseStatuscode(t, statusCodeUpdateArticle, http.StatusOK)
	})
}

func TestReadArticle(t *testing.T) {
	var serveHTTP routesFunc
	serveHTTP = controllers.NewWikiHandlers().ArticleRoutes

	newArticle := "A wiki is a knowledge base website"
	ioReader := strings.NewReader(newArticle)
	reqPutArticle := newPutArticleRequest("wiki", ioReader)
	resPutArticle := httptest.NewRecorder()
	serveHTTP(resPutArticle, reqPutArticle)

	t.Run("it returns 200 OK when the article is found", func(t *testing.T) {
		req := newGetArticleRequest("wiki")
		res := httptest.NewRecorder()
		serveHTTP(res, req)

		gotContentType := res.Result().Header.Get("content-type")
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		content := string(bodyBytes)

		if content != newArticle {
			t.Errorf("expected content %v; but got %v", newArticle, content)
		}
		assertContentType(t, gotContentType, "text/html")
	})

	t.Run("it returns 404 Not Found when the article is not found", func(t *testing.T) {
		req := newGetArticleRequest("rest_api")
		res := httptest.NewRecorder()
		serveHTTP(res, req)

		statusCode := res.Result().StatusCode
		assertResponseStatuscode(t, statusCode, http.StatusNotFound)
	})

	t.Run("it returns an array that has an element", func(t *testing.T) {
		req := newGetArticleRequest("")
		res := httptest.NewRecorder()
		serveHTTP(res, req)

		gotContentType := res.Result().Header.Get("content-type")
		assertContentType(t, gotContentType, "application/json")

		got, _ := json.Marshal(res.Body)

		if len(got) == 0 {
			t.Errorf("expected an array that has elements but got an empty array")
		}
	})
}

func newGetArticleRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/articles/%s", name), nil)
	return req
}

func newPutArticleRequest(name string, httpBody *strings.Reader) *http.Request {
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/articles/%s", name), httpBody)
	return req
}

func assertResponseStatuscode(t *testing.T, gotStatusCode, expectedStatusCode int) {
	t.Helper()
	if gotStatusCode != expectedStatusCode {
		t.Errorf("expected status %v; but got %v", expectedStatusCode, gotStatusCode)
	}
}

func assertResponsePayload(t *testing.T, gotPayload, msg string) {
	t.Helper()
	if len(gotPayload) > 0 {
		t.Errorf(msg, gotPayload)
	}
}

func assertContentType(t *testing.T, gotContentType, expectedContentType string) {
	t.Helper()

	if gotContentType != expectedContentType {
		t.Errorf("expected contentType application/json; but got %v", gotContentType)
	}
}

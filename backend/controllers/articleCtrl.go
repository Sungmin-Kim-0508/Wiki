package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/wiki/backend/myutils"
)

// Wiki type
type Wiki struct {
	Name    string
	Content string
}

type wikiHandlers struct {
	sync.Mutex
	store map[string]string
}

// NewWikiHandlers returns an instance of wikiHandler object
func NewWikiHandlers() *wikiHandlers {
	return &wikiHandlers{
		store: map[string]string{},
	}
}

func (w *wikiHandlers) ArticleRoutes(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		splitURL := myutils.SplitURLBySlash(req.URL.String())
		parsedURL := myutils.SliceNil(splitURL)
		developmentURLLength := 2
		if len(parsedURL) == developmentURLLength {
			w.getSingleArticle(res, req)
			return
		}
		w.getArticlesList(res, req)
		return
	case "PUT":
		w.editArticle(res, req)
		return
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte("Method Not Allowed"))
		return
	}
}

// GET all aricles
func (w *wikiHandlers) getArticlesList(res http.ResponseWriter, req *http.Request) {
	wikis := make([]Wiki, len(w.store))
	w.Lock()
	i := 0
	for name, content := range w.store {
		wikis[i] = Wiki{Name: name, Content: content}
		i++
	}
	w.Unlock()

	jsonBytes, err := json.Marshal(wikis)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
	}
	res.Header().Add("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
	return
}

// GET single article
func (w *wikiHandlers) getSingleArticle(res http.ResponseWriter, req *http.Request) {
	parsedURL := strings.Split(req.URL.String(), "/")
	if len(parsedURL) != 3 {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	name := parsedURL[2]
	article, exists := w.store[name]
	if exists {
		res.Header().Add("content-type", "text/html")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(article))
		return
	}

	res.WriteHeader(http.StatusNotFound)
	return
}

// PUT create or edit article
func (w *wikiHandlers) editArticle(res http.ResponseWriter, req *http.Request) {
	parsedURL := strings.Split(req.URL.String(), "/")
	if len(parsedURL) != 3 {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	name := parsedURL[2]
	content := string(bodyBytes)

	_, exists := w.store[name]

	// Update Article
	if exists {
		res.WriteHeader(http.StatusOK)
		w.store[name] = content
		return
	}
	// Create Article
	w.store[name] = content
	res.WriteHeader(http.StatusCreated)
	return
}

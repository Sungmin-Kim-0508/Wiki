package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/wiki/backend/utils"
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
	baseURL := utils.GetBaseURL()
	switch req.Method {
	case "GET":
		name := strings.TrimPrefix(req.URL.Path, baseURL+"articles/")
		if len(name) > 0 {
			w.getSingleArticle(res, req, name)
			return
		}
		w.getArticlesList(res, req)
		return
	case "PUT":
		w.mutateArticle(res, req)
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
func (w *wikiHandlers) getSingleArticle(res http.ResponseWriter, req *http.Request, name string) {
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
func (w *wikiHandlers) mutateArticle(res http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	baseURL := utils.GetBaseURL()
	name := strings.TrimPrefix(req.URL.Path, baseURL+"articles/")
	content := string(bodyBytes)

	_, exists := w.store[name]

	// Update Article
	if exists {
		res.WriteHeader(http.StatusOK)
		w.store[name] = content
		return
	}
	// Create Article
	res.WriteHeader(http.StatusCreated)
	w.store[name] = content
	return
}

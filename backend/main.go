package main

import (
	"fmt"
	"net/http"

	"github.com/wiki/backend/controllers"
	"github.com/wiki/backend/utils"
)

func CORS(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func main() {
	wikis := controllers.NewWikiHandlers()
	baseURL := utils.GetBaseURL()
	http.HandleFunc(baseURL, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	http.HandleFunc(baseURL+"articles/", CORS(wikis.ArticleRoutes))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}

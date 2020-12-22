package main

import (
	"fmt"
	"net/http"

	"github.com/wiki/backend/controllers"
)

func main() {
	wikis := controllers.NewWikiHandlers()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.HandleFunc("/articles/", wikis.ArticleRoutes)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/aseure/jagger/query"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := query.NewQuery(r.Body)

	if query == nil {
		w.WriteHeader(400)
	} else {
		go query.Execute()
		w.WriteHeader(200)
	}
}

func main() {
	fmt.Println("Listening on :8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

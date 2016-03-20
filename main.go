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
		response := make(chan string)
		w.WriteHeader(200)
		go query.Execute(response)
		fmt.Fprintf(w, "%s", <-response)
	}
}

func main() {
	fmt.Println("Listening on :8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

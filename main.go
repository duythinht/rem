package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	R := mux.NewRouter()
	R.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.ListenAndServe(":8000", R)
}

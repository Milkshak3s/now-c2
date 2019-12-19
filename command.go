package main

import (
	"fmt"
	"net/http"
)

func Command(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.RawQuery)
}

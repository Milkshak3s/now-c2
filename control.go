package main

import (
	"fmt"
	"net/http"
)

func Control(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.RawQuery)
}

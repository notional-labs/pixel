package controller

import (
	"fmt"
	"net/http"
)

func GetPixelHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/pixels" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// to do add get pixels func
	fmt.Fprintf(w, "Hello!")
}

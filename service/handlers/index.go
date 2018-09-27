package handlers

import (
	"fmt"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No content on %s\n", html.EscapeString(r.URL.Path))
}
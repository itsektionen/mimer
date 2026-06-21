package app

import "net/http"

func HxRefresh(w http.ResponseWriter) {
	w.Header().Set("HX-Refresh", "true")
}

func HxRedirect(w http.ResponseWriter, location string) {
	w.Header().Set("HX-Redirect", location)
}

func HxLocation(w http.ResponseWriter, location string) {
	w.Header().Set("HX-Location", location)
}

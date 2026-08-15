package app

import "net/http"

func HXRefresh(w http.ResponseWriter) {
	w.Header().Set("HX-Refresh", "true")
}

func HXRedirect(w http.ResponseWriter, location string) {
	w.Header().Set("HX-Redirect", location)
}

func HXLocation(w http.ResponseWriter, location string) {
	w.Header().Set("HX-Location", location)
}

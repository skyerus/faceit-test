package api

import "net/http"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, 200, nil)
}

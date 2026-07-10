package main

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := HealthResponse{Status: "ok"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// TODO: que faire si l'encodage échoue ?
	}
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func routeHandler(proxy *httputil.ReverseProxy, targets map[string]url.URL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := targets[r.Host]
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			// TODO: change page not found message
			resp := ErrorResponse{Code: "404", Message: "Page not found."}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				// TODO: que faire si l'encodage échoue ?
			}
			return
		}
		proxy.ServeHTTP(w, r)
	}
}

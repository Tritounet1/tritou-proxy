package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	health(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected application/json got %q", ct)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	var healthResponse HealthResponse
	if err := json.Unmarshal(data, &healthResponse); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if healthResponse.Status != "ok" {
		t.Fatalf("expected ok got %v", healthResponse.Status)
	}
}

func TestRouteHandlerUnknowHost(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	targets := make(map[string]url.URL)
	targets["hote-inconnu.test"] = url.URL{Scheme: "http", Host: "hote-inconnu.test", Path: "/"}
	proxy := buildProxy(false, targets)
	handler := routeHandler(&proxy, targets)
	handler(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 got %d", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected application/json got %q", ct)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	var errorResponse ErrorResponse
	if err := json.Unmarshal(data, &errorResponse); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if errorResponse.Code != "404" {
		t.Fatalf("expected 404 got %v", errorResponse.Code)
	}
}

func TestRouteHandlerKnowHost(t *testing.T) {
	expected := "dummy data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "dummy data")
	}))
	defer svr.Close()
	urlParse, err := url.Parse(svr.URL)
	if err != nil {
		t.Fatalf("bad url: %v", err)
	}
	targetHost := "backend.test"
	targets := make(map[string]url.URL)
	targets[targetHost] = *urlParse
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Host = targetHost
	w := httptest.NewRecorder()
	proxy := buildProxy(false, targets)
	handler := routeHandler(&proxy, targets)
	handler(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	got := strings.TrimSpace(string(data))
	if got != expected {
		t.Errorf("expected %q got %q", expected, got)
	}
}

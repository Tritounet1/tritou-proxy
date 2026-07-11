package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestProxyRewritingHost(t *testing.T) {
	var gotReq *http.Request
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotReq = r
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
	req.RemoteAddr = "203.0.113.5:1234"
	w := httptest.NewRecorder()
	proxy := buildProxy(false, targets)
	handler := routeHandler(&proxy, targets)
	handler(w, req)
	if gotReq.Host != "backend.test" {
		t.Errorf("expected backend.test host got %s", gotReq.Host)
	}
	if gotReq.Header.Get("X-Forwarded-For") != "203.0.113.5" {
		t.Errorf("expected 203.0.113.5 forwaded for got %s", gotReq.Header.Get("X-Forwarded-For"))
	}
	if gotReq.Header.Get("X-Forwarded-Host") != "backend.test" {
		t.Errorf("expected backend.test forwaded host got %s", gotReq.Header.Get("X-Forwarded-Host"))
	}
	if gotReq.Header.Get("X-Forwarded-Proto") != "http" {
		t.Errorf("expected http forwaded proto got %s", gotReq.Header.Get("X-Forwarded-Proto"))
	}
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", res.StatusCode)
	}
}

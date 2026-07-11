package main

import (
	"testing"
)

func TestBuildTargets(t *testing.T) {
	cfg := Config{
		Routes: map[string]Route{
			"backend.test": {
				Target: Target{
					Scheme: "http",
					Host:   "localhost:3000",
					Path:   "/",
				},
			},
		},
	}
	targets := buildTargets(cfg)
	if len(targets) == 0 {
		t.Fatalf("expected 1 items in targets, got %d", len(targets))
	}
	target := targets["backend.test"]
	if target.Scheme != "http" {
		t.Fatalf("expected http scheme, got %q", target.Scheme)
	}
	if target.Host != "localhost:3000" {
		t.Fatalf("expected localhost:3000 host, got %q", target.Host)
	}
	if target.Path != "/" {
		t.Fatalf("expected '/' path, got %q", target.Path)
	}
}

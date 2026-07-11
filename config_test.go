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
	targets, err := buildTargets(cfg)
	if err != nil {
		t.Fatalf("unexpected an error for invalid config")
	}
	if len(targets) != 1 {
		t.Fatalf("expected 1 item in targets, got %d", len(targets))
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

func TestLoadYamlFile(t *testing.T) {
	// Test correct yaml
	correct_cfg, err := loadConfig("testdata/correct.yaml")
	if err != nil {
		t.Fatalf("error while trying to load config.yaml file %q", err)
	}
	correct_targets, err := buildTargets(correct_cfg)
	if err != nil {
		t.Fatalf("error while trying to build targets %q", err)
	}
	if len(correct_targets) != 1 {
		t.Fatalf("expected 1 item in targets, got %d", len(correct_targets))
	}
	correct_target := correct_targets["backend.test"]
	if correct_target.Scheme != "http" {
		t.Fatalf("expected http scheme, got %q", correct_target.Scheme)
	}
	if correct_target.Host != "localhost:3000" {
		t.Fatalf("expected localhost:3000 host, got %q", correct_target.Host)
	}
	if correct_target.Path != "/" {
		t.Fatalf("expected '/' path, got %q", correct_target.Path)
	}
	// Test incorrect yaml
	_, err = loadConfig("testdata/incorrect.yaml")
	if err == nil {
		t.Fatalf("excepted error while trying to load config.yaml file")
	}
}

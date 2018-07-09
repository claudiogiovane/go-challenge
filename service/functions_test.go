package main

import (
	"testing"
)

// Tests valid and invalid URLs via 'validURL'.
func TestValidURL(t *testing.T) {
	URLs := []struct {
		valid   string
		invalid string
	}{
		{"https://10.65.10.188", "htp://10.65.10.188"},
		{"http://localhost:8080", "localhost::8080"},
		{"http://www.google.com", "http:/www.google.com"},
		{"https://site.org", "httpss://testurl.com"},
		{"http://www.teste.de", "http//www.teste.de"},
	}

	for _, url := range URLs {
		ok := validURL(url.valid)
		nok := validURL(url.invalid)
		switch {
		case ok != nil:
			t.Errorf("URL %s must return OK", url.valid)
		case nok == nil:
			t.Errorf("URL %s must return an error", url.invalid)
		}
	}
}

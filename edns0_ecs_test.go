package edns0_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CumpsD/edns0"
)

func TestDemo(t *testing.T) {
	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	cfg := edns0.CreateConfig()
	cfg.Prefix = "TESTPREFIX"

	handler, err := edns0.New(ctx, next, cfg, "edns0")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)
}

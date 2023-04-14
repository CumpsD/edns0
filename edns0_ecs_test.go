package edns0_ecs_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	edns0_ecs "github.com/CumpsD/traefik-edns0-ecs-middleware"
)

func TestDemo(t *testing.T) {
	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	cfg := edns0_ecs.CreateConfig()
	cfg.Prefix = "TESTPREFIX"

	handler, err := edns0_ecs.New(ctx, next, cfg, "edns0-ecs")
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

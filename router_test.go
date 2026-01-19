package simplerouter

import (
	"net/http"
	"testing"
)

func TestNew_WithServerMux_UsesProvidedMux(t *testing.T) {
	mux := http.NewServeMux()
	r := New(WithServerMux(mux))

	if r.mux != mux {
		t.Fatalf("expected provided to be used, got different instance")
	}
}

func TestNew(t *testing.T) {
	r := New()

	r.Get("/hello", func(c *Ctx) error {
		return nil
	})
}

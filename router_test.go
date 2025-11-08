package httprouter_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jeremyctrl/httprouter"
)

func TestRouteMatch(t *testing.T) {
	router := httprouter.New().
		GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			io.WriteString(w, "ok")
		}).
		Build()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if got := w.Body.String(); got != "ok" {
		t.Fatalf("expected 'ok', got %q", got)
	}
}

func TestParamsExtraction(t *testing.T) {
	router := httprouter.New().
		GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			io.WriteString(w, ps.ByName("name"))
		}).
		Build()

	req := httptest.NewRequest("GET", "/hello/alice", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if got := strings.TrimSpace(w.Body.String()); got != "alice" {
		t.Fatalf("expected 'alice', got %q", got)
	}
}

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New().
		GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}).
		Build()

	router.NotFound = func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "custom 404", http.StatusNotFound)
	}

	req := httptest.NewRequest("GET", "/missing", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if !strings.Contains(w.Body.String(), "custom 404") {
		t.Fatalf("expected custom 404 response, got %q", w.Body.String())
	}
}

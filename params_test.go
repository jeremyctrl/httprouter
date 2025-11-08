package httprouter_test

import (
	"testing"

	"github.com/jeremyctrl/httprouter"
)

func TestParamsByName(t *testing.T) {
	ps := httprouter.Params{
		{Name: "user", Value: "alice"},
		{Name: "id", Value: "42"},
	}

	if got := ps.ByName("user"); got != "alice" {
		t.Errorf("expected 'alice', got %q", got)
	}

	if got := ps.ByName("id"); got != "42" {
		t.Errorf("expected '42', got %q", got)
	}

	if got := ps.ByName("missing"); got != "" {
		t.Errorf("expected empty string for missing param, got %q", got)
	}
}

func TestParamsByPosition(t *testing.T) {
	ps := httprouter.Params{
		{Name: "first", Value: "one"},
		{Name: "second", Value: "two"},
	}

	if got := ps.ByPosition(0); got != "one" {
		t.Errorf("expected 'one', got %q", got)
	}
	if got := ps.ByPosition(1); got != "two" {
		t.Errorf("expected 'two', got %q", got)
	}
	if got := ps.ByPosition(2); got != "" {
		t.Errorf("expected empty string for out of range, got %q", got)
	}
}

func TestParamsNamesAndValues(t *testing.T) {
	ps := httprouter.Params{
		{Name: "a", Value: "x"},
		{Name: "b", Value: "y"},
		{Name: "c", Value: "z"},
	}

	names := ps.Names()
	values := ps.Values()

	expectedNames := []string{"a", "b", "c"}
	expectedValues := []string{"x", "y", "z"}

	for i := range expectedNames {
		if names[i] != expectedNames[i] {
			t.Errorf("expected name %q at position %d, got %q", expectedNames[i], i, names[i])
		}
		if values[i] != expectedValues[i] {
			t.Errorf("expected value %q at position %d, got %q", expectedValues[i], i, values[i])
		}
	}
}

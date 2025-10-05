package random

import (
	"slices"
	"testing"
)

func TestChoose_EmptySlice(t *testing.T) {
	rs := New[int]()
	var items []int

	got := rs.Choose(items)
	var zero int
	if got != zero {
		t.Errorf("expected zero value, got %v", got)
	}
}

func TestChoose_SingleElement(t *testing.T) {
	rs := New[string]()
	want := "hello"
	items := []string{want}

	got := rs.Choose(items)
	if got != want {
		t.Errorf("expected %q, got %v", want, got)
	}
}

func TestChoose_MultipleElements(t *testing.T) {
	rs := New[int]()
	items := []int{1, 2, 3, 4, 5}

	got := rs.Choose(items)
	if !slices.Contains(items, got) {
		t.Errorf("element %v not found in items %v", got, items)
	}
}

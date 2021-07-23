package main

import "testing"

func TestSlice(t *testing.T) {
	got := MakeSlice(5)
	want := []string{"1", "2", "3", "4", "5"}

	if len(got) != len(want) {
		t.Errorf("got length %d want %d", len(got), len(want))
	}

}

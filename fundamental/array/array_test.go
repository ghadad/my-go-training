package main

import "testing"

func TestArray(t *testing.T) {
	got := Initilze()
	want := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	want[1] = 2
	if got == want {
		t.Errorf("got %q want %q", got, want)
	}

	g := GetArr2()

	if g != [3]int{1, 2, 3} {
		t.Errorf("got %q want %q", g, want)
	}

}

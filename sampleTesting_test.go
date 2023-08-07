package main

import "testing"

func TestAddSum(t *testing.T) {
	got := addSum(1, 4)
	want := 5

	if (got != want) {
		t.Errorf("Got %q, Wanted %q", got, want)
	}
}
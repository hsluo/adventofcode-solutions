package main

import "testing"

func TestWrap(t *testing.T) {
	if wrap([]int{2, 3, 4}) != 58 {
		t.Fail()
	}
	if wrap([]int{1, 1, 10}) != 43 {
		t.Fail()
	}
}

func TestRibbon(t *testing.T) {
	if ribbon([]int{2, 3, 4}) != 34 {
		t.Fail()
	}
	if ribbon([]int{1, 1, 10}) != 14 {
		t.Fail()
	}
}

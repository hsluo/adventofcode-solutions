package main

import "testing"

func TestCalc(t *testing.T) {
	if calc("2x3x4") != 58 {
		t.Fail()
	}
	if calc("1x1x10") != 43 {
		t.Fail()
	}
}

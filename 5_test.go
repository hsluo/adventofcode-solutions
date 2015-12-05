package main

import "testing"

func TestNice(t *testing.T) {
	nices := []string{"ugknbfddgicrmopn", "aaa"}
	for i := range nices {
		if !nice(nices[i]) {
			t.Log(nices[i])
			t.Fail()
		}
	}

	naughties := []string{"jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb"}
	for i := range naughties {
		if nice(naughties[i]) {
			t.Log(naughties[i])
			t.Fail()
		}
	}
}

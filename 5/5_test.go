package main

import "testing"

func TestNice1(t *testing.T) {
	nices := []string{"ugknbfddgicrmopn", "aaa"}
	for i := range nices {
		if !nice1(nices[i]) {
			t.Log(nices[i])
			t.Fail()
		}
	}

	naughties := []string{"jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb"}
	for i := range naughties {
		if nice1(naughties[i]) {
			t.Log(naughties[i])
			t.Fail()
		}
	}
}

func TestNice2(t *testing.T) {
	nices := []string{"qjhvhtzxzqqjkmpb", "xxyxx"}
	for i := range nices {
		if !nice2(nices[i]) {
			t.Log(nices[i])
			t.Fail()
		}
	}
	naughties := []string{"uurcxstgmygtbstg", "ieodomkazucvgmuy"}
	for i := range naughties {
		if nice1(naughties[i]) {
			t.Log(naughties[i])
			t.Fail()
		}
	}
}

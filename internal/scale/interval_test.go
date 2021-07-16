package scale

import "testing"

func TestInterval_IsEquivalent(t *testing.T) {
	a := Intervals.Natural9
	b := Intervals.Major2

	if !a.IsEquivalent(&b) {
		t.Error("Expected: 2 and 9 is Equivalent")
	}

	c := Intervals.R
	d := Intervals.Major2

	if c.IsEquivalent(&d) {
		t.Error("Expected: R and 2 is not Equivalent")
	}
}

func TestInterval_IsEquals(t *testing.T) {
	a := Intervals.Natural9
	b := Intervals.Major2

	if a.IsEquals(&b) {
		t.Error("Expected: 2 and 9 is not Equals")
	}

	c := Intervals.Natural9
	if !a.IsEquals(&c) {
		t.Error("Expected: 9 and 9 is Equals")
	}
}
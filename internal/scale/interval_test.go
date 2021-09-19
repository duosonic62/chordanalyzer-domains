package scale

import "testing"

func TestNewInterval(t *testing.T) {
	interval, _ := NewInterval("R")
	if !interval.Equals(&Intervals.R) {
		t.Error("Expected: R, but actual: " + interval.String())
	}

	interval, _ = NewInterval("b9")
	if !interval.Equals(&Intervals.Flat9) {
		t.Error("Expected: b9, but actual: " + interval.String())
	}

	_, err := NewInterval("m9")
	if err == nil {
		t.Error("Expected: error is not nil, but error is nil")
	}
}

func TestInterval_IsEquivalent(t *testing.T) {
	a := Intervals.Natural9
	b := Intervals.Major2

	if !a.Equivalent(&b) {
		t.Error("Expected: 2 and 9 is Equivalent")
	}

	c := Intervals.R
	d := Intervals.Major2

	if c.Equivalent(&d) {
		t.Error("Expected: R and 2 is not Equivalent")
	}
}

func TestInterval_IsEquals(t *testing.T) {
	a := Intervals.Natural9
	b := Intervals.Major2

	if a.Equals(&b) {
		t.Error("Expected: 2 and 9 is not Equals")
	}

	c := Intervals.Natural9
	if !a.Equals(&c) {
		t.Error("Expected: 9 and 9 is Equals")
	}
}

func TestInterval_intervalFromNumber(t *testing.T) {
	actual, _ := intervalFromNumber(0)
	if !Intervals.R.Equivalent(actual) {
		t.Error("Expected: R, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(1)
	if !Intervals.Minor2.Equivalent(actual) {
		t.Error("Expected: Minor2, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(2)
	if !Intervals.Major2.Equivalent(actual) {
		t.Error("Expected: Major2, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(3)
	if !Intervals.Minor3.Equivalent(actual) {
		t.Error("Expected: Minor3, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(4)
	if !Intervals.Major3.Equivalent(actual) {
		t.Error("Expected: Major3, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(5)
	if !Intervals.Perfect4.Equivalent(actual) {
		t.Error("Expected: Perfect4, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(6)
	if !Intervals.Sharp4.Equivalent(actual) {
		t.Error("Expected: Sharp4, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(7)
	if !Intervals.Perfect5.Equivalent(actual) {
		t.Error("Expected: Perfect5, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(8)
	if !Intervals.Minor6.Equivalent(actual) {
		t.Error("Expected: Minor6, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(9)
	if !Intervals.Major6.Equivalent(actual) {
		t.Error("Expected: Major6, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(10)
	if !Intervals.Minor7.Equivalent(actual) {
		t.Error("Expected: Minor7, but actual: " + actual.name)
	}

	actual, _ = intervalFromNumber(11)
	if !Intervals.Major7.Equivalent(actual) {
		t.Error("Expected: Major7, but actual: " + actual.name)
	}
}
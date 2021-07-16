package tone

import "testing"

func TestString_PositiveWholeTone(t *testing.T) {
	actual, _ := NewInterval(10)
	if actual.String() != "#6" || actual.String() != "m7" {
		t.Error("Expected: #6 or m7, but actual: " + actual.String())
	}
}

func TestString_PositiveHalfTone(t *testing.T) {
	actual, _ := NewInterval(11)
	if actual.String() != "M7" {
		t.Error("Expected: 5.5, but actual: " + actual.String())
	}
}

func TestInterval_EquivalentIsEquivalent(t *testing.T) {
	// #6とb7は等価
	a, _ := NewIntervalFromName("#6")
	b, _ := NewIntervalFromName("m7")
	if !a.IsEquivalent(b) {
		t.Error("Expected: #6 and m7 is Equivalent")
	}
}

func TestInterval_NotEquivalentIsEquivalent(t *testing.T) {
	// M7とm7は等価じゃない
	a, _ := NewIntervalFromName("M7")
	b, _ := NewIntervalFromName("m7")
	if a.IsEquivalent(b) {
		t.Error("Expected: M6 and m7 is not Equivalent")
	}
}
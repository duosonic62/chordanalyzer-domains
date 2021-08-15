package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"testing"
)

func TestFactory_Build_Triad(t *testing.T) {
	intervals := []scale.Interval{
		scale.Intervals.R,
		scale.Intervals.Major3,
		scale.Intervals.Perfect5,
	}

	factory := NewCodeFactory(intervals)
	codes, _ := factory.Build()
	if "C" != codes[0].Name() {
		t.Error("Expected: C, but actual: " + codes[0].Name())
	}
	if "C#" != codes[1].Name() {
		t.Error("Expected: C#, but actual: " + codes[1].Name())
	}
	if "D" != codes[2].Name() {
		t.Error("Expected: D, but actual: " + codes[2].Name())
	}
	if "D#" != codes[3].Name() {
		t.Error("Expected: D#, but actual: " + codes[3].Name())
	}
	if "E" != codes[4].Name() {
		t.Error("Expected: E, but actual: " + codes[4].Name())
	}
	if "F" != codes[5].Name() {
		t.Error("Expected: F, but actual: " + codes[5].Name())
	}
	if "F#" != codes[6].Name() {
		t.Error("Expected: F#, but actual: " + codes[6].Name())
	}
	if "G" != codes[7].Name() {
		t.Error("Expected: G, but actual: " + codes[7].Name())
	}
	if "G#" != codes[8].Name() {
		t.Error("Expected: G#, but actual: " + codes[8].Name())
	}
	if "A" != codes[9].Name() {
		t.Error("Expected: A, but actual: " + codes[9].Name())
	}
	if "A#" != codes[10].Name() {
		t.Error("Expected: A#, but actual: " + codes[10].Name())
	}
	if "B" != codes[11].Name() {
		t.Error("Expected: B, but actual: " + codes[11].Name())
	}
}

func TestFactory_Build_Tension(t *testing.T) {
	intervals := []scale.Interval{
		scale.Intervals.R,
		scale.Intervals.Major3,
		scale.Intervals.Perfect5,
		scale.Intervals.Minor7,
	}

	factory := NewCodeFactory(intervals)
	codes, _ := factory.Build()
	if "C7" != codes[0].Name() {
		t.Error("Expected: C7, but actual: " + codes[0].Name())
	}
	if "C#7" != codes[1].Name() {
		t.Error("Expected: C#7, but actual: " + codes[1].Name())
	}
	if "D7" != codes[2].Name() {
		t.Error("Expected: D7, but actual: " + codes[2].Name())
	}
	if "D#7" != codes[3].Name() {
		t.Error("Expected: D#7, but actual: " + codes[3].Name())
	}
	if "E7" != codes[4].Name() {
		t.Error("Expected: E7, but actual: " + codes[4].Name())
	}
	if "F7" != codes[5].Name() {
		t.Error("Expected: F7, but actual: " + codes[5].Name())
	}
	if "F#7" != codes[6].Name() {
		t.Error("Expected: F#7, but actual: " + codes[6].Name())
	}
	if "G7" != codes[7].Name() {
		t.Error("Expected: G7, but actual: " + codes[7].Name())
	}
	if "G#7" != codes[8].Name() {
		t.Error("Expected: G#7, but actual: " + codes[8].Name())
	}
	if "A7" != codes[9].Name() {
		t.Error("Expected: A7, but actual: " + codes[9].Name())
	}
	if "A#7" != codes[10].Name() {
		t.Error("Expected: A#7, but actual: " + codes[10].Name())
	}
	if "B7" != codes[11].Name() {
		t.Error("Expected: B7, but actual: " + codes[11].Name())
	}
}

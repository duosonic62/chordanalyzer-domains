package chord

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
	"testing"
)

func TestFactory_Build(t *testing.T) {
	factory := NewCodeFactory(Major, []scale.Interval{scale.Intervals.Minor7})
	codes, err := factory.Build()
	if err != nil {
		t.Error("Unexpected error: " + err.Error())
	}
	if len(codes) != 12 {
		t.Error("Expected: 12 chords, but actual: " + strconv.Itoa(len(codes)))
	}

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

func TestFactory_BuildError(t *testing.T) {
	factory := NewCodeFactory(Major, []scale.Interval{})
	_, err := factory.Build()
	if err == nil {
		t.Error("Expected error occurred,  but actual error is null")
	}
}

func TestFactory_BuildWithName(t *testing.T) {
	factory := NewCodeFactory(Diminish, []scale.Interval{scale.Intervals.Minor7})
	codes, err := factory.BuildWithName("m7b5")
	if err != nil {
		t.Error("Unexpected error: " + err.Error())
	}
	if len(codes) != 12 {
		t.Error("Expected: 12 chords, but actual: " + strconv.Itoa(len(codes)))
	}

	if "Cm7b5" != codes[0].Name() {
		t.Error("Expected: Cm7b5, but actual: " + codes[0].Name())
	}
	if "C#m7b5" != codes[1].Name() {
		t.Error("Expected: C#m7b5, but actual: " + codes[1].Name())
	}
	if "Dm7b5" != codes[2].Name() {
		t.Error("Expected: Dm7b5, but actual: " + codes[2].Name())
	}
	if "D#m7b5" != codes[3].Name() {
		t.Error("Expected: D#m7b5, but actual: " + codes[3].Name())
	}
	if "Em7b5" != codes[4].Name() {
		t.Error("Expected: Em7b5, but actual: " + codes[4].Name())
	}
	if "Fm7b5" != codes[5].Name() {
		t.Error("Expected: Fm7b5, but actual: " + codes[5].Name())
	}
	if "F#m7b5" != codes[6].Name() {
		t.Error("Expected: F#m7b5, but actual: " + codes[6].Name())
	}
	if "Gm7b5" != codes[7].Name() {
		t.Error("Expected: Gm7b5, but actual: " + codes[7].Name())
	}
	if "G#m7b5" != codes[8].Name() {
		t.Error("Expected: G#m7b5, but actual: " + codes[8].Name())
	}
	if "Am7b5" != codes[9].Name() {
		t.Error("Expected: Am7b5, but actual: " + codes[9].Name())
	}
	if "A#m7b5" != codes[10].Name() {
		t.Error("Expected: A#m7b5, but actual: " + codes[10].Name())
	}
	if "Bm7b5" != codes[11].Name() {
		t.Error("Expected: Bm7b5, but actual: " + codes[11].Name())
	}
}

func TestFactory_BuildWithNameError(t *testing.T) {
	factory := NewCodeFactory(Diminish, []scale.Interval{})
	_, err := factory.BuildWithName("m7b5")
	if err == nil {
		t.Error("Expected error occurred,  but actual error is null")
	}
}

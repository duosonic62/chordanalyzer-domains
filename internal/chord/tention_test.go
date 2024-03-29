package chord

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"testing"
)

func TestNewTensionChord(t *testing.T) {
	actual, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major6})
	if actual.Name() != "C6" {
		t.Error("Expected: C6, but actual: " + actual.Name())
	}

	// under 3 notes
	_, err := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestNewTensionChordFrom(t *testing.T) {
	actual, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major6})
	if actual.Name() != "C6" {
		t.Error("Expected: C6, but actual: " + actual.Name())
	}

	// under 3 notes
	_, err := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestNewTensionChordWithNameFrom(t *testing.T) {
	actual, _ := NewTensionChordWithName("m7b5", &scale.Notes.C, Diminish, []scale.Interval{scale.Intervals.Major7})
	if actual.Name() != "Cm7b5" {
		t.Error("Expected: Cm7b5, but actual: " + actual.Name())
	}

	// under 3 notes
	_, err := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestTension_Name(t *testing.T) {
	actual, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major6})
	if actual.Name() != "C6" {
		t.Error("Expected: C6, but actual: " + actual.Name())
	}
}

func TestTension_Root(t *testing.T) {
	actual, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major6})
	if !scale.Notes.C.Equals(actual.Root()) {
		t.Error("Expected: C, but actual: " + actual.Root().String())
	}
}

func TestTension_Notes(t *testing.T) {
	actual, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major6})
	if !actual.Notes()[0].Equals(&scale.Notes.C) && actual.Notes()[1].Equals(&scale.Notes.E) && actual.Notes()[2].Equals(&scale.Notes.G) && actual.Notes()[3].Equals(&scale.Notes.A) {
		t.Error("Expected: C, E, G, A, but actual: " + actual.Notes()[0].String() + ", " + actual.Notes()[1].String() + ", " + actual.Notes()[2].String() + ", " + actual.Notes()[3].String())
	}
}

func TestTension_Intervals(t *testing.T) {
	actual, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major6})
	if !actual.Intervals()[0].Equals(&scale.Intervals.R) && actual.Intervals()[1].Equals(&scale.Intervals.Major3) && actual.Intervals()[2].Equals(&scale.Intervals.Perfect5) && actual.Intervals()[3].Equals(&scale.Intervals.Major6) {
		t.Error("Expected: C, E, G, A, but actual: " + actual.Intervals()[0].String() + ", " + actual.Intervals()[1].String() + ", " + actual.Intervals()[2].String() + ", " + actual.Intervals()[3].String())
	}
}

func TestTension_Contains(t *testing.T) {
	c7, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major7})
	c, _ := NewTriad(&scale.Notes.C, Major)
	em, _ := NewTriad(&scale.Notes.E, Minor)

	if !c7.Contains(c) {
		t.Error("Expected: CM7 contains C, but actual doesn't contain")
	}

	if !c7.Contains(em) {
		t.Error("Expected: CM7 contains Em, but actual doesn't contain")
	}

	if em.Contains(c7) {
		t.Error("Expected: Em doesn't contains CM7, but actual　contains")
	}
}

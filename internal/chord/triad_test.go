package chord

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"testing"
)


func TestNewTriadFrom(t *testing.T) {
	// Major Chord
	actual, _ := NewTriad(&scale.Notes.C, Major)
	if actual.Name() != "C" {
		t.Error("Expected: C, but actual: " + actual.Name())
	}

	// Minor
	actual, _ = NewTriad(&scale.Notes.C, Minor)
	if actual.Name() != "Cm" {
		t.Error("Expected: Cm, but actual: " + actual.Name())
	}

	// Aug
	actual, _ = NewTriad(&scale.Notes.C, Augment)
	if actual.Name() != "Caug" {
		t.Error("Expected: Caug, but actual: " + actual.Name())
	}

	// Dim
	actual, _ = NewTriad(&scale.Notes.C, Diminish)
	if actual.Name() != "Cdim" {
		t.Error("Expected: Cdim, but actual: " + actual.Name())
	}

	// Major b5
	actual, _ = NewTriad(&scale.Notes.C, MajorB5)
	if actual.Name() != "CMb5" {
		t.Error("Expected: Cdim, but actual: " + actual.Name())
	}

	// sus2
	actual, _ = NewTriad(&scale.Notes.C, Sus2)
	if actual.Name() != "Csus2" {
		t.Error("Expected: Csus2, but actual: " + actual.Name())
	}

	// sus4
	actual, _ = NewTriad(&scale.Notes.C, Sus4)
	if actual.Name() != "Csus4" {
		t.Error("Expected: Csus4, but actual: " + actual.Name())
	}

	// no Triad
	_, err := NewTriad(&scale.Notes.C, TriadType("sample"))
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestTriad_Name(t *testing.T) {
	// name == Name()
	actual, _ := NewTriad(&scale.Notes.C, Major)
	if actual.Name() != "C" {
		t.Error("Expected: " + actual.Name() + ", but actual: C")
	}
}

func TestTriad_Root(t *testing.T) {
	// Root == notes[0]
	actual, _ := NewTriad(&scale.Notes.C, Major)
	if !scale.Notes.C.Equals(actual.Root()) {
		t.Error("Expected: C, but actual: " + actual.Name())
	}
}

func TestTriad_Notes(t *testing.T) {
	// name == Name()
	actual, _ := NewTriad(&scale.Notes.C, Major)
	if !actual.Notes()[0].Equals(&scale.Notes.C) && actual.Notes()[1].Equals(&scale.Notes.E) && actual.Notes()[2].Equals(&scale.Notes.G) {
		t.Error("Expected: " + actual.Name() + ", but actual: " + actual.Name())
	}
}

func TestTriad_Intervals(t *testing.T) {
	actual, _ := NewTriad(&scale.Notes.C, Major)
	if !actual.Intervals()[0].Equals(&scale.Intervals.R) && actual.Intervals()[1].Equals(&scale.Intervals.Major3) && actual.Intervals()[2].Equals(&scale.Intervals.Perfect5) {
		t.Error("Expected: C, E, G, but actual: " + actual.Intervals()[0].String() + ", " + actual.Intervals()[1].String() + ", " + actual.Intervals()[2].String())
	}
}

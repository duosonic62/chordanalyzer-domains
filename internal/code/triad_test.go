package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"testing"
)

func TestNewTriad(t *testing.T) {
	// Major Code
	actual, _ := NewTriad([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G})
	if actual.name != "C" {
		t.Error("Expected: C, but actual: " + actual.name)
	}

	// Minor
	actual, _ = NewTriad([]scale.Note{scale.Notes.C, scale.Notes.Eb, scale.Notes.G})
	if actual.name != "Cm" {
		t.Error("Expected: Cm, but actual: " + actual.name)
	}

	// Aug
	actual, _ = NewTriad([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.GS})
	if actual.name != "Caug" {
		t.Error("Expected: Caug, but actual: " + actual.name)
	}

	// Dim
	actual, _ = NewTriad([]scale.Note{scale.Notes.C, scale.Notes.Eb, scale.Notes.Gb})
	if actual.name != "Cdim" {
		t.Error("Expected: Cdim, but actual: " + actual.name)
	}

	// no 3 notes
	_, err := NewTriad([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G, scale.Notes.B})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}

	// no triads
	_, err = NewTriad([]scale.Note{scale.Notes.C, scale.Notes.G, scale.Notes.B})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestTriad_Name(t *testing.T) {
	// name == Name()
	actual, _ := NewTriad([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G})
	if actual.name != actual.Name() {
		t.Error("Expected: " + actual.name + ", but actual: " + actual.Name())
	}
}

func TestTriad_Root(t *testing.T) {
	// Root == notes[0]
	actual, _ := NewTriad([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G})
	if !scale.Notes.C.Equals(actual.Root()) {
		t.Error("Expected: " + actual.name + ", but actual: " + actual.Name())
	}
}

func TestTriad_Notes(t *testing.T) {
	// name == Name()
	actual, _ := NewTriad([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G})
	if !actual.Notes()[0].Equals(&scale.Notes.C) && actual.Notes()[1].Equals(&scale.Notes.E) && actual.Notes()[2].Equals(&scale.Notes.G)  {
		t.Error("Expected: " + actual.name + ", but actual: " + actual.Name())
	}
}
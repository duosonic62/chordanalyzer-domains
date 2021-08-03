package scale

import "testing"

func TestNewTriad(t *testing.T) {
	// Major Code
	actual, _ := NewTriad([]Note{Notes.C, Notes.E, Notes.G})
	if actual.name != "C" {
		t.Error("Expected: C, but actual: " + actual.name)
	}

	// Minor
	actual, _ = NewTriad([]Note{Notes.C, Notes.Eb, Notes.G})
	if actual.name != "Cm" {
		t.Error("Expected: Cm, but actual: " + actual.name)
	}

	// Aug
	actual, _ = NewTriad([]Note{Notes.C, Notes.E, Notes.GS})
	if actual.name != "Caug" {
		t.Error("Expected: Caug, but actual: " + actual.name)
	}

	// Dim
	actual, _ = NewTriad([]Note{Notes.C, Notes.Eb, Notes.Gb})
	if actual.name != "Cdim" {
		t.Error("Expected: Cdim, but actual: " + actual.name)
	}

	// no 3 notes
	_, err := NewTriad([]Note{Notes.C, Notes.E, Notes.G, Notes.B})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}

	// no triads
	_, err = NewTriad([]Note{Notes.C, Notes.G, Notes.B})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestTriad_Name(t *testing.T) {
	// name == Name()
	actual, _ := NewTriad([]Note{Notes.C, Notes.E, Notes.G})
	if actual.name != actual.Name() {
		t.Error("Expected: " + actual.name + ", but actual: " + actual.Name())
	}
}

func TestTriad_Root(t *testing.T) {
	// Root == notes[0]
	actual, _ := NewTriad([]Note{Notes.C, Notes.E, Notes.G})
	if !Notes.C.Equals(actual.Root()) {
		t.Error("Expected: " + actual.name + ", but actual: " + actual.Name())
	}
}

func TestTriad_Notes(t *testing.T) {
	// name == Name()
	actual, _ := NewTriad([]Note{Notes.C, Notes.E, Notes.G})
	if !actual.Notes()[0].Equals(&Notes.C) && actual.Notes()[1].Equals(&Notes.E) && actual.Notes()[2].Equals(&Notes.G)  {
		t.Error("Expected: " + actual.name + ", but actual: " + actual.Name())
	}
}
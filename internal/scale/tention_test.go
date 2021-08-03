package scale

import "testing"

func TestNewTensionCode(t *testing.T) {
	actual, _ := NewTensionCode([]Note{Notes.C, Notes.E, Notes.G, Notes.A})
	if actual.name != "C6" {
		t.Error("Expected: C6, but actual: " + actual.name)
	}

	// under 3 notes
	_, err := NewTensionCode([]Note{Notes.C, Notes.E, Notes.G})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestTension_Name(t *testing.T) {
	actual, _ := NewTensionCode([]Note{Notes.C, Notes.E, Notes.G, Notes.A})
	if actual.Name() != "C6" {
		t.Error("Expected: C6, but actual: " + actual.name)
	}
}

func TestTension_Root(t *testing.T) {
	actual, _ := NewTensionCode([]Note{Notes.C, Notes.E, Notes.G, Notes.A})
	if !Notes.C.Equals(actual.Root()) {
		t.Error("Expected: C, but actual: " + actual.Root().name)
	}
}

func TestTension_Notes(t *testing.T) {
	actual, _ := NewTensionCode([]Note{Notes.C, Notes.E, Notes.G, Notes.A})
	if !actual.Notes()[0].Equals(&Notes.C) && actual.Notes()[1].Equals(&Notes.E) && actual.Notes()[2].Equals(&Notes.G) && actual.Notes()[3].Equals(&Notes.A) {
		t.Error("Expected: C, E, G, A, but actual: " + actual.Notes()[0].name + ", " + actual.Notes()[1].name + ", " + actual.Notes()[2].name + ", " + actual.Notes()[3].name)
	}
}

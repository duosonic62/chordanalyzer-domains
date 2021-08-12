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

func TestTension_Contains(t *testing.T) {
	c7, _ := NewTensionCode([]Note{Notes.C, Notes.E, Notes.G, Notes.B})
	c, _ := NewTriad([]Note{Notes.C, Notes.E, Notes.G})
	em, _ := NewTriad([]Note{Notes.E, Notes.G, Notes.B})

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
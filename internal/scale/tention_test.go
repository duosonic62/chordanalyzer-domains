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
package scale

import "testing"

func TestNewTriad(t *testing.T) {
	// Major Code
	actual, _ := NewTriad([]Note{Notes.C, Notes.E, Notes.G})
	if actual.name != "C" {
		t.Error("Expected: Major3, but actual: " + actual.name)
	}

	// Minor
	actual, _ = NewTriad([]Note{Notes.C, Notes.DS, Notes.G})
	if actual.name != "C" {
		t.Error("Expected: Major3, but actual: " + actual.name)
	}
}
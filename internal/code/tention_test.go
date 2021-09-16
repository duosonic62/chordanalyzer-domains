package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"testing"
)

func TestNewTensionCode(t *testing.T) {
	actual, _ := NewTensionCode([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G, scale.Notes.A})
	if actual.name != "C6" {
		t.Error("Expected: C6, but actual: " + actual.name)
	}

	// under 3 notes
	_, err := NewTensionCode([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G})
	if err == nil {
		t.Error("Expected: err, but error is Nil")
	}
}

func TestTension_Name(t *testing.T) {
	actual, _ := NewTensionCode([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G, scale.Notes.A})
	if actual.Name() != "C6" {
		t.Error("Expected: C6, but actual: " + actual.name)
	}
}

func TestTension_Root(t *testing.T) {
	actual, _ := NewTensionCode([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G, scale.Notes.A})
	if !scale.Notes.C.Equals(actual.Root()) {
		t.Error("Expected: C, but actual: " + actual.Root().String())
	}
}

func TestTension_Notes(t *testing.T) {
	actual, _ := NewTensionCode([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G, scale.Notes.A})
	if !actual.Notes()[0].Equals(&scale.Notes.C) && actual.Notes()[1].Equals(&scale.Notes.E) && actual.Notes()[2].Equals(&scale.Notes.G) && actual.Notes()[3].Equals(&scale.Notes.A) {
		t.Error("Expected: C, E, G, A, but actual: " + actual.Notes()[0].String() + ", " + actual.Notes()[1].String() + ", " + actual.Notes()[2].String() + ", " + actual.Notes()[3].String())
	}
}

func TestTension_Contains(t *testing.T) {
	c7, _ := NewTensionCode([]scale.Note{scale.Notes.C, scale.Notes.E, scale.Notes.G, scale.Notes.B})
	c, _ := NewTriad(&scale.Notes.C, []scale.Interval{scale.Intervals.R, scale.Intervals.Major3, scale.Intervals.Perfect5})
	em, _ := NewTriad(&scale.Notes.E, []scale.Interval{scale.Intervals.R, scale.Intervals.Minor3, scale.Intervals.Perfect5})

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
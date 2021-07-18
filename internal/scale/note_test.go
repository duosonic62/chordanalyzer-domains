package scale

import (
	"testing"
)

func TestNote_CalculateInterval(t *testing.T) {
	// 基音 < 比較音
	actual, _ := Notes.C.CalculateInterval(Notes.D)
	if !actual.IsEquivalent(&Intervals.Major2) {
		t.Error("Expected: Major 2, but actual: " + actual.name)
	}

	// 基音 > 比較音
	actual, _ = Notes.G.CalculateInterval(Notes.C)
	if !actual.IsEquivalent(&Intervals.Perfect4) {
		t.Error("Expected: Major 3, but actual: " + actual.name)
	}

	// 基音 = 比較音
	actual, _ = Notes.G.CalculateInterval(Notes.G)
	if !actual.IsEquivalent(&Intervals.R) {
		t.Error("Expected: R, but actual: " + actual.name)
	}
}

func TestNote_GetIntervalNote(t *testing.T) {
	actual, _ := Notes.C.GetIntervalNote(&Intervals.Perfect5)
	if actual != &Notes.G {
		t.Error("Expected: G, but actual: " + actual.name)
	}

	actual, _ = Notes.G.GetIntervalNote(&Intervals.Perfect4)
	if actual != &Notes.C {
		t.Error("Expected: C, but actual: " + actual.name)
	}
}


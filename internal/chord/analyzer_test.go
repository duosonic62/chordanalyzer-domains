package chord

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
	"testing"
)

func TestAnalyzer_AnalyzeIncludedChords(t *testing.T) {
	analyzer := NewAnalyzer(collections)
	cM7, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major7})
	actual := analyzer.AnalyzeIncludedChords(cM7)
	if len(actual) != 2 {
		t.Error("Expected 2 chords, but actual " + strconv.Itoa(len(actual)))
	}

	if actual[0].Name() != "C" && actual[1].Name() != "C" && actual[2].Name() != "C" {
		t.Error("Expected contains C, but actual not contains")
	}

	if actual[0].Name() != "Em" && actual[1].Name() != "Em" && actual[2].Name() != "Em" {
		t.Error("Expected contains Em, but actual not contains")
	}
}

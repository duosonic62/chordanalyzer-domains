package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
	"testing"
)

func TestAnalyzer_AnalyzeIncludedCodes(t *testing.T) {
	analyzer := NewAnalyzer(collections)
	cM7, _ := NewTensionCode([]scale.Note{
		scale.Notes.C,
		scale.Notes.E,
		scale.Notes.G,
		scale.Notes.B,
	})
	actual := analyzer.AnalyzeIncludedCodes(cM7)
	if len(actual) != 3 {
		t.Error("Expected 3 codes, but actual " + strconv.Itoa(len(actual)))
	}

	if actual[0].Name() != "C" && actual[1].Name() != "C" && actual[2].Name() != "C" {
		t.Error("Expected contains C, but actual not contains")
	}

	if actual[0].Name() != "Em" && actual[1].Name() != "Em" && actual[2].Name() != "Em" {
		t.Error("Expected contains Em, but actual not contains")
	}

	if actual[0].Name() != "CM7" && actual[1].Name() != "CM7" && actual[2].Name() != "CM7" {
		t.Error("Expected contains CM7, but actual not contains")
	}
}
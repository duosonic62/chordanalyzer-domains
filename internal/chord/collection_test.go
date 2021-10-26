package chord

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
	"strings"
	"testing"
)

func TestCollection_Filter(t *testing.T) {
	actual := major7Collection.Filter(func(code Chord) bool {
		return strings.Contains(code.Name(), "C")
	})

	if len(actual.ToSlice()) != 16 {
		t.Error("Expected: C* and C#* but not")
	}
}

func TestCollection_ForEach(t *testing.T) {
	count := 0
	major7Collection.ForEach(func(code Chord) {
		count++
	})

	if count/len(scale.AllNotes()) != len(AllTriadTypes)+1 {
		t.Error("Expected: " + strconv.Itoa(len(AllTriadTypes)+1) + " actions but actual: " + strconv.Itoa(count/len(scale.AllNotes())))
	}
}

func TestCollection_Get(t *testing.T) {
	actual := major7Collection.Get("CM7")
	major7, _ := NewTensionCode(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major7})
	if actual.Name() != major7.Name() {
		t.Error("")
	}
}

func TestCollection_GetNoCodes(t *testing.T) {
	actual := major7Collection.Get("Cm7")
	if actual != nil {
		t.Error("")
	}
}


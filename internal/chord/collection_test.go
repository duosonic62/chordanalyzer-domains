package chord

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"strconv"
	"strings"
	"testing"
)

func TestCollection_Filter(t *testing.T) {
	actual := major7Collection.Filter(func(chord Chord) bool {
		return strings.Contains(chord.Name(), "C")
	})

	if len(actual.ToSlice()) != 16 {
		t.Error("Expected: C* and C#* but not")
	}
}

func TestCollection_ForEach(t *testing.T) {
	count := 0
	major7Collection.ForEach(func(chord Chord) {
		count++
	})

	if count/len(scale.AllNotes()) != len(AllTriadTypes)+1 {
		t.Error("Expected: " + strconv.Itoa(len(AllTriadTypes)+1) + " actions but actual: " + strconv.Itoa(count/len(scale.AllNotes())))
	}
}

func TestCollection_Get(t *testing.T) {
	actual := major7Collection.Get("CM7")
	major7, _ := NewTensionChord(&scale.Notes.C, Major, []scale.Interval{scale.Intervals.Major7})
	if actual.Name() != major7.Name() {
		t.Error("Expected: " + major7.Name() + " , but actual: " + actual.Name())
	}
}

func TestCollection_GetNoChords(t *testing.T) {
	actual := major7Collection.Get("Cm7")
	if actual != nil {
		t.Error("")
	}
}


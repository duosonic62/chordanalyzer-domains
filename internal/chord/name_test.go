package chord

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"testing"
)

func TestName_Get(t *testing.T) {
	actual := NewName(&scale.Notes.Db, Major, []scale.Interval{scale.Intervals.Minor7, scale.Intervals.Flat9}).Get()

	if actual != "Db7b9" {
		t.Error("Expected: Db7b9, but actual: " + actual)
	}
}

func TestName_Equals(t *testing.T) {
	actual := NewName(&scale.Notes.Db, Major, []scale.Interval{scale.Intervals.Minor7, scale.Intervals.Flat9})

	if !actual.Equals("Db7b9") {
		t.Error("Db7b9 expected to equal Db7b9, but actual is false")
	}

	if !actual.Equals("C#7b9") {
		t.Error("Db7b9 expected to equal C#7b9, but actual is false")
	}
}

func Test_NewName(t *testing.T) {
	actual := NewName(&scale.Notes.Db, Major, []scale.Interval{scale.Intervals.Minor7, scale.Intervals.Flat9}).Get()
	if actual != "Db7b9" {
		t.Error("Expected: Db7b9, but actual: " + actual)
	}

	actual = NewName(&scale.Notes.Db, Minor, []scale.Interval{scale.Intervals.Minor7, scale.Intervals.Flat9}).Get()
	if actual != "Dbm7b9" {
		t.Error("Expected: Dbm7b9, but actual: " + actual)
	}

	actual = NewName(&scale.Notes.Db, Sus4, []scale.Interval{scale.Intervals.Minor7, scale.Intervals.Flat9}).Get()
	if actual != "Dbsus47b9" {
		t.Error("Expected: Dbsus47b9, but actual: " + actual)
	}
}

func Test_NewNameWithTensionName(t *testing.T) {
	actual := NewNameWithTensionName(&scale.Notes.Db, "7b5").Get()

	if actual != "Db7b5" {
		t.Error("Expected: Db7b9, but actual: " + actual)
	}
}
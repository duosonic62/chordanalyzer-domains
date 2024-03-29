package chord

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Tension struct {
	name             Name
	triad            *Triad
	tensionNotes     []scale.Note
	tensionIntervals []scale.Interval
}

func (t Tension) Name() string {
	return t.name.Get()
}

func (t Tension) Root() *scale.Note {
	return t.triad.Root()
}

func (t Tension) Notes() []scale.Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]scale.Note, t.notesNum())
	copy(notes, append(t.triad.Notes(), t.tensionNotes...))
	return notes
}

func (t Tension) Intervals() []scale.Interval {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	intervals := make([]scale.Interval, t.notesNum())
	copy(intervals, append(t.triad.Intervals(), t.tensionIntervals...))
	return intervals
}

func (t Tension) Contains(other Chord) bool {
	// otherの方がノート数が多ければfalse
	if len(other.Notes()) > len(t.Notes()) {
		return false
	}

	// 1ノートでも含まれていないノートがあれば false
	for _, on := range other.Notes() {
		if !t.contains(&on) {
			return false
		}
	}

	return true
}

func (t Tension) CompareByName(name string) bool {
	return t.name.Equals(name)
}

func (t Tension) contains(note *scale.Note) bool {
	for _, tn := range t.Notes() {
		if tn.Equivalent(note) {
			return true
		}
	}

	return false
}

func (t Tension) notesNum() int {
	return len(t.tensionNotes) + len(t.triad.notes)
}

func NewTensionChord(root *scale.Note, triadType TriadType, tensions []scale.Interval) (*Tension, error) {
	if len(tensions) < 1 {
		return nil, errors.New("value chord must contains over 4 notes")
	}

	var tensionName string
	for _, interval := range tensions {
		tensionName = tensionName + interval.String()
	}

	tensionNotes, err := intervalsToNotes(root, tensions)
	if err != nil {
		return nil, err
	}

	triad, err := NewTriad(root, triadType)
	if err != nil {
		return nil, err
	}

	return &Tension{
		name:         NewName(root, triadType, tensions),
		triad:        triad,
		tensionNotes: tensionNotes,
		tensionIntervals: tensions,
	}, nil
}

func NewTensionChordWithName(name string, root *scale.Note, triadType TriadType, tensions []scale.Interval) (*Tension, error) {
	if len(tensions) < 1 {
		return nil, errors.New("value chord must contains over 4 notes")
	}

	tensionNotes, err := intervalsToNotes(root, tensions)
	if err != nil {
		return nil, err
	}

	triad, err := NewTriad(root, triadType)
	if err != nil {
		return nil, err
	}

	return &Tension{
		name:         NewNameWithTensionName(root, name),
		triad:        triad,
		tensionNotes: tensionNotes,
		tensionIntervals: tensions,
	}, nil
}

func intervalsToNotes(root *scale.Note, intervals []scale.Interval) ([]scale.Note, error) {
	notes := make([]scale.Note, len(intervals))
	for i, interval := range intervals {
		note, err := root.GetIntervalNote(&interval)
		if err != nil {
			return nil, err
		}
		notes[i] = *note
	}

	return notes, nil
}

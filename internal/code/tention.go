package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Tension struct {
	name             string
	triad            *Triad
	tensionNotes     []scale.Note
	tensionIntervals []scale.Interval
}

func (t Tension) Name() string {
	return t.name
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

func (t Tension) Contains(other Code) bool {
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

func NewTensionCode(root *scale.Note, intervals []scale.Interval) (*Tension, error) {
	if len(intervals) < 4 {
		return nil, errors.New("tension code must contains over 4 notes")
	}
	triad, err := NewTriad(root, intervals[0:3])
	if err != nil {
		return nil, errors.Wrap(err, "tension code must contains triad")
	}

	tensionIntervals := intervals[3:]

	var tensionName string
	for _, interval := range tensionIntervals {
		tensionName = tensionName + interval.String()
	}

	tensionNotes, err := intervalsToNotes(root, tensionIntervals)
	if err != nil {
		return nil, err
	}

	return &Tension{
		name:         triad.name + tensionName,
		triad:        triad,
		tensionNotes: tensionNotes,
		tensionIntervals: tensionIntervals,
	}, nil
}

func NewTensionCodeFrom(root *scale.Note, triadType TriadType, tensions []scale.Interval) (*Tension, error) {
	if len(tensions) < 1 {
		return nil, errors.New("tension code must contains over 4 notes")
	}

	var tensionName string
	for _, interval := range tensions {
		tensionName = tensionName + interval.String()
	}

	tensionNotes, err := intervalsToNotes(root, tensions)
	if err != nil {
		return nil, err
	}

	triad, err := NewTriadFrom(root, triadType)
	if err != nil {
		return nil, err
	}

	return &Tension{
		name:         triad.name + tensionName,
		triad:        triad,
		tensionNotes: tensionNotes,
		tensionIntervals: tensions,
	}, nil
}

func NewTensionCodeWithName(name string, root *scale.Note, intervals []scale.Interval) (*Tension, error) {
	if len(intervals) < 4 {
		return nil, errors.New("tension code must contains over 4 notes")
	}
	triad, err := NewTriad(root, intervals[0:3])
	if err != nil {
		return nil, errors.Wrap(err, "tension code must contains triad")
	}

	tensionIntervals := intervals[3:]
	tensionNotes, err := intervalsToNotes(root, tensionIntervals)
	if err != nil {
		return nil, err
	}

	return &Tension{
		name:         name,
		triad:        triad,
		tensionNotes: tensionNotes,
		tensionIntervals: tensionIntervals,
	}, nil
}

func NewTensionCodeWithNameFrom(name string, root *scale.Note, triadType TriadType, tensions []scale.Interval) (*Tension, error) {
	if len(tensions) < 1 {
		return nil, errors.New("tension code must contains over 4 notes")
	}

	tensionNotes, err := intervalsToNotes(root, tensions)
	if err != nil {
		return nil, err
	}

	triad, err := NewTriadFrom(root, triadType)
	if err != nil {
		return nil, err
	}

	return &Tension{
		name:         name,
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

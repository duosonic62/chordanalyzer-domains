package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Tension struct {
	name  string
	triad *Triad
	tensionNotes []scale.Note
}

func (t Tension) Name() string {
	return t.name
}

func (t Tension) Root() *scale.Note {
	return t.triad.root
}

func (t Tension) Notes() []scale.Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]scale.Note, t.notesNum())
	copy(notes, append(t.triad.Notes(), t.tensionNotes...))
	return notes
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

func NewTensionCode(notes []scale.Note) (*Tension, error) {
	if len(notes) < 4 {
		return nil, errors.New("tension code must contains over 4 notes")
	}
	triad, err := NewTriad(notes[0:3])
	if err != nil {
		return nil, errors.Wrap(err, "tension code must contains triad")
	}

	tensionNotes := notes[3:]

	var tensionName string
	for _, note := range tensionNotes {
		interval, err := triad.root.CalculateTensionInterval(note)
		if err != nil {
			return nil, errors.Wrap(err, "contains invalid intervals")
		}
		tensionName = tensionName + interval.String()
	}
	return &Tension{
		name:         triad.name + tensionName,
		triad:        triad,
		tensionNotes: tensionNotes,
	}, nil
}
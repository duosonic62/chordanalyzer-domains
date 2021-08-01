package scale

import "github.com/pkg/errors"

type Tension struct {
	name  string
	triad *Triad
	tensionNotes []Note
}

func (t Tension) Name() string {
	return t.name
}

func (t Tension) Root() *Note {
	return t.triad.root
}

func (t Tension) Notes() []Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]Note, t.notesNum())
	copy(notes, append(t.triad.Notes(), t.tensionNotes...))
	return notes
}

func (t Tension) notesNum() int {
	return len(t.tensionNotes) + len(t.triad.notes)
}

func NewTensionCode(notes []Note) (*Tension, error) {
	triad, err := NewTriad(notes[0:2])
	if err != nil {
		return nil, errors.Wrap(err, "tension code must contains triad")
	}

	tensionNotes := notes[3:]

	var tensionName string
	for _, note := range tensionNotes {
		interval, err := triad.root.CalculateInterval(note)
		if err != nil {
			return nil, errors.Wrap(err, "contains invalid intervals")
		}
		tensionName = tensionName + interval.name
	}
	return &Tension{
		name:         triad.name + tensionName,
		triad:        triad,
		tensionNotes: tensionNotes,
	}, nil
}
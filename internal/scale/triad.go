package scale

import "github.com/pkg/errors"

type Triad struct {
	name  string
	root  *Note
	notes []Note
}

func (t Triad) Name() string {
	return t.name
}

func (t Triad) Root() *Note {
	return t.root
}

func (t Triad) Notes() []Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]Note, len(t.notes), cap(t.notes))
	copy(notes, t.notes)
	return notes
}

//NewTriad トライアドを生成する
func NewTriad(notes []Note) (*Triad, error) {
	if len(notes) != 3 {
		return nil, errors.New("the number of notes in the triad must be 3")
	}

	root := notes[0]
	third, err := root.CalculateInterval(notes[1])
	fifth, err := root.CalculateInterval(notes[2])
	if err != nil {
		return nil, errors.Wrap(err, "contains invalid intervals")
	}

	// major triad
	if third.IsEquivalent(&Intervals.Major3) && fifth.IsEquals(&Intervals.Perfect5) {
		return &Triad{
			name:  root.name,
			root:  &root,
			notes: notes,
		}, nil
	}
	// minor triad
	if third.IsEquivalent(&Intervals.Minor3) && fifth.IsEquals(&Intervals.Perfect5) {
		return &Triad{
			name:  root.name + "m",
			root:  &root,
			notes: notes,
		}, nil
	}
	// augmented triad
	if third.IsEquivalent(&Intervals.Major3) && fifth.IsEquals(&Intervals.Sharp5) {
		return &Triad{
			name:  root.name + "aug",
			root:  &root,
			notes: notes,
		}, nil
	}
	// diminished triad
	if third.IsEquivalent(&Intervals.Minor3) && fifth.IsEquals(&Intervals.Flat5) {
		return &Triad{
			name:  root.name + "dim",
			root:  &root,
			notes: notes,
		}, nil
	}

	return nil, errors.New("triads must be [major, minor, aug, dim]")
}

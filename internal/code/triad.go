package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Triad struct {
	name  string
	root  *scale.Note
	notes []scale.Note
}

func (t Triad) Name() string {
	return t.name
}

func (t Triad) Root() *scale.Note {
	return t.root
}

func (t Triad) Notes() []scale.Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]scale.Note, len(t.notes), cap(t.notes))
	copy(notes, t.notes)
	return notes
}

func (t Triad) Contains(other Code) bool {
	// トライアドが別のコードを含むことはないので、名前が一致した時のみtrue
	return t.Name() == other.Name()
}

//NewTriad トライアドを生成する
func NewTriad(root *scale.Note, intervals []scale.Interval) (*Triad, error) {
	if len(intervals) != 3 {
		return nil, errors.New("the number of notes in the triad must be 3")
	}

	third := intervals[1]
	fifth := intervals[2]

	return newTriad(root, []scale.Interval{scale.Intervals.R, third, fifth})
}

func newTriad(root *scale.Note, intervals []scale.Interval) (*Triad, error) {
	if len(intervals) != 3 {
		return nil, errors.New("the number of notes in the triad must be 3")
	}

	third := intervals[1]
	fifth := intervals[2]
	notes := make([]scale.Note, 3, 3)

	for i, interval := range intervals {
		note, err := root.GetIntervalNote(&interval)
		if err != nil {
			return nil, err
		}
		notes[i] = *note
	}

	// major triad
	if third.IsEquivalent(&scale.Intervals.Major3) && fifth.IsEquals(&scale.Intervals.Perfect5) {
		return &Triad{
			name:  root.String(),
			root:  root,
			notes: notes,
		}, nil
	}
	// minor triad
	if third.IsEquivalent(&scale.Intervals.Minor3) && fifth.IsEquals(&scale.Intervals.Perfect5) {
		return &Triad{
			name:  root.String() + "m",
			root:  root,
			notes: notes,
		}, nil
	}
	// augmented triad
	if third.IsEquivalent(&scale.Intervals.Major3) && fifth.IsEquals(&scale.Intervals.Sharp5) {
		return &Triad{
			name:  root.String() + "aug",
			root:  root,
			notes: notes,
		}, nil
	}
	// diminished triad
	if third.IsEquivalent(&scale.Intervals.Minor3) && fifth.IsEquals(&scale.Intervals.Flat5) {
		return &Triad{
			name:  root.String() + "dim",
			root:  root,
			notes: notes,
		}, nil
	}

	return nil, errors.New("triads must be [major, minor, aug, dim]")
}
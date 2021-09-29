package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Triad struct {
	name      string
	root      scale.Note
	notes     []scale.Note
	intervals []scale.Interval
}

func (t Triad) Name() string {
	return t.name
}

func (t Triad) Root() *scale.Note {
	return &t.root
}

func (t Triad) Notes() []scale.Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]scale.Note, len(t.notes), cap(t.notes))
	copy(notes, t.notes)
	return notes
}

func (t Triad) Intervals() []scale.Interval {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	intervals := make([]scale.Interval, len(t.notes), cap(t.notes))
	copy(intervals, t.intervals)
	return intervals
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
	notes := make([]scale.Note, 3)

	for i, interval := range intervals {
		note, err := root.GetIntervalNote(&interval)
		if err != nil {
			return nil, err
		}
		notes[i] = *note
	}

	// major triad
	if third.Equivalent(&scale.Intervals.Major3) && fifth.Equals(&scale.Intervals.Perfect5) {
		return &Triad{
			name:  root.String(),
			root:  *root,
			notes: notes,
		}, nil
	}
	// minor triad
	if third.Equivalent(&scale.Intervals.Minor3) && fifth.Equals(&scale.Intervals.Perfect5) {
		return &Triad{
			name:  root.String() + "m",
			root:  *root,
			notes: notes,
		}, nil
	}
	// augmented triad
	if third.Equivalent(&scale.Intervals.Major3) && fifth.Equals(&scale.Intervals.Sharp5) {
		return &Triad{
			name:  root.String() + "aug",
			root:  *root,
			notes: notes,
		}, nil
	}
	// diminished triad
	if third.Equivalent(&scale.Intervals.Minor3) && fifth.Equals(&scale.Intervals.Flat5) {
		return &Triad{
			name:  root.String() + "dim",
			root:  *root,
			notes: notes,
		}, nil
	}

	return nil, errors.New("triads must be [major, minor, aug, dim]")
}

type TriadEnum string

const (
	Major    TriadEnum = ""
	Minor    TriadEnum = "m"
	Augment  TriadEnum = "aug"
	Diminish TriadEnum = "dim"
	MajorB5  TriadEnum = "b5"
	Sus2     TriadEnum = "sus2"
	Sus4     TriadEnum = "sus4"
)

func NewTriadFrom(root *scale.Note, enum TriadEnum) (*Triad, error) {
	third, err :=  root.GetIntervalNote(&scale.Intervals.Major3)
	fifth, err := root.GetIntervalNote(&scale.Intervals.Perfect5)
	if err != nil {
		return nil, err
	}

	return &Triad{
		name:      root.String() + string(enum),
		root:      *root,
		notes:     []scale.Note{*root, *third, *fifth},
		intervals: []scale.Interval{scale.Intervals.R, scale.Intervals.Major3, scale.Intervals.Perfect5},
	}, nil
}
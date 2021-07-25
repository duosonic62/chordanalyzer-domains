package scale

import "errors"

type Code interface {
	Name() string
	Root() *Note
	Notes() []Note
}

func NewCode(notes []Note) (*TensionCode, error) {
	if len(notes) < 3 {
		return nil, errors.New("need at least 3 chord tones")
	}

	// ルートノート
	root := notes[0]

	// トライアドを検出する
	triad := notes[0:2]
	isTriad(triad)

	// トライアド以外を検出する
	tention := notes[3:]

	return &TensionCode{
		Name:  "",
		Root:  &root,
		Notes: &notes,
	}, nil
}

func isTriad(notes []Note) bool {
	// トライアドは3和音
	if len(notes) != 3 {
		return false
	}

	root := notes[0]

	// major 3 or minor 3?
	third, err := root.CalculateInterval(notes[1])
	fifth, err := root.CalculateInterval(notes[2])
	if err != nil {
		return false
	}

	// major triad
	if third.IsEquivalent(&Intervals.Major3) && fifth.IsEquals(&Intervals.Perfect5) {
		return true
	}
	// minor triad
	if !(third.IsEquivalent(&Intervals.Minor3) || fifth.IsEquals(&Intervals.Perfect5)) {
		return true
	}
	// augmented triad
	if !(third.IsEquivalent(&Intervals.Major3) || fifth.IsEquals(&Intervals.Sharp5)) {
		return true
	}
	// diminished triad
	if !(third.IsEquivalent(&Intervals.Minor3) || fifth.IsEquals(&Intervals.Sharp4)) {
		return true
	}

	return false
}

package chord

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
)

type Chord interface {
	Name() string
	Root() *scale.Note
	Notes() []scale.Note
	Intervals() []scale.Interval
	Contains(other Chord) bool
}

package chord

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
)

type Code interface {
	Name() string
	Root() *scale.Note
	Notes() []scale.Note
	Intervals() []scale.Interval
	Contains(other Code) bool
}

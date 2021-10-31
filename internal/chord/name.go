package chord

import "github.com/duosonic62/chordanalyzer-domains/internal/scale"

type Name interface {
	Get() string
	Equals() bool
}

type name struct {
	root scale.Note
	tension string
}
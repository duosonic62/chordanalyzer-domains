package chord

import "github.com/duosonic62/chordanalyzer-domains/internal/scale"

type Name interface {
	Get() string
	Equals(name string) bool
}

type name struct {
	root scale.Note
	tension string
}

func (n name) Get() string {
	return n.root.String() + n.tension
}

func (n name) Equals(name string) bool {
	return name == n.Get()
}

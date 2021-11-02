package chord

import "github.com/duosonic62/chordanalyzer-domains/internal/scale"

type Name interface {
	Get() string
	Equals(name string) bool
}

type name struct {
	value string
}

func (n name) Get() string {
	return n.value
}

func (n name) Equals(name string) bool {
	return name == n.Get()
}

func NewName(root *scale.Note, tensions []scale.Interval) Name {
	var tensionName string
	for _, interval := range tensions {
		tensionName = tensionName + interval.String()
	}

	return name{
		value: root.String() + tensionName,
	}
}

func NewNameWithTensionName(root *scale.Note, tension string) Name {
	return name{
		root.String() + tension,
	}
}

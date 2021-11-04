package chord

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"regexp"
)

type Name interface {
	Get() string
	Equals(name string) bool
}

var codeNameRegex = regexp.MustCompile("([A-G][b#]?)(.*)")

type name struct {
	root    scale.Note
	tension string
	value   string
}

func (n name) Get() string {
	return n.root.String() + n.tension
}

func (n name) Equals(name string) bool {
	codeName := codeNameRegex.FindStringSubmatch(name)
	if len(codeName) != 2 {
		return false
	}
	rootName := codeName[0]
	tensionName := codeName[1]
	// root0] を scale.Noteに変換

	// root[0] と n.root, n.tension と tensionを比較
	// return root.Equivalent(n.root) && tensionName == n.tension
	return name == n.Get()
}

func NewName(root *scale.Note, tensions []scale.Interval) Name {
	var tensionName string
	for _, interval := range tensions {
		tensionName = tensionName + interval.String()
	}

	return name{
		root:  *root,
		tension: tensionName,
		value: root.String() + tensionName,
	}
}

func NewNameWithTensionName(root *scale.Note, tension string) Name {
	return name{
		root:  *root,
		tension: tension,
		value: root.String() + tension,
	}
}

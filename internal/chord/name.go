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
	root  scale.Note
	value string
}

func (n name) Get() string {
	return n.root.String() + n.value
}

func (n name) Equals(name string) bool {
	codeName := codeNameRegex.FindStringSubmatch(name)
	if len(codeName) != 3 {
		return false
	}
	rootName := codeName[1]
	tensionName := codeName[2]

	// root0] を scale.Noteに変換
	root, err := scale.FromString(rootName)
	if err != nil {
		return false
	}

	// root[0] と n.root, n.value と tensionを比較
	return root.Equivalent(&n.root) && tensionName == n.value
}

func NewName(root *scale.Note, triad TriadType, tensions []scale.Interval) Name {
	var tensionName string
	for _, interval := range tensions {
		tensionName = tensionName + interval.String()
	}

	return name{
		root:  *root,
		value: string(triad) + tensionName,
	}
}

func NewNameWithTensionName(root *scale.Note, tension string) Name {
	return name{
		root:  *root,
		value: tension,
	}
}

package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Factory struct {
	triad   TriadType
	tension []scale.Interval
}

func NewCodeFactory(triad TriadType, tension []scale.Interval) *Factory {
	return &Factory{
		triad:   triad,
		tension: tension,
	}
}

//Build is create code instances slice
func (f Factory) Build() ([]Code, error) {
	if len(f.tension) == 0 {
		return nil, errors.New("failed to build codes, there are no intervals")
	}

	tensions, err := f.createTensionCode()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build codes, there are incorrect intervals")
	}
	return tensions, nil
}

//BuildWithName is create tension code with name
//Use this function to generate code whose name is difficult to guess. (e.g. Cm7b5 and Cdim7)
//For triad codes, the name will be ignored.
func (f Factory) BuildWithName(name string) ([]Code, error) {
	if len(f.tension) == 0 {
		return nil, errors.New("failed to build codes, there are no intervals")
	}

	tensions, err := f.createTensionWithName(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build codes, there are incorrect intervals")
	}
	return tensions, nil
}

func (f Factory) createTensionCode() ([]Code, error) {
	allNotes := scale.AllNotes()
	tensions := make([]Code, len(allNotes))
	for i, root := range scale.AllNotes() {
		tension, err := NewTensionCodeFrom(&root, f.triad, f.tension)
		if err != nil {
			return nil, err
		}
		tensions[i] = tension
	}

	return tensions, nil
}

func (f Factory) createTensionWithName(name string) ([]Code, error) {
	allNotes := scale.AllNotes()
	tensions := make([]Code, len(allNotes))
	for i, root := range scale.AllNotes() {
		tension, err := NewTensionCodeWithNameFrom(root.String() + name, &root, f.triad, f.tension)
		if err != nil {
			return nil, err
		}
		tensions[i] = tension
	}

	return tensions, nil
}

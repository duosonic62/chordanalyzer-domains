package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Code interface {
	Name() string
	Root() *scale.Note
	Notes() []scale.Note
	Contains(other Code) bool
}

type Factory struct {
	intervals []scale.Interval
}

func NewFactory(intervals []scale.Interval) *Factory {
	return &Factory{
		intervals: intervals,
	}
}

//Build is create code instances slice
func (f Factory) Build() ([]Code, error) {

	if len(f.intervals) == 3 {
		triads, err := f.createTriads()
		if err != nil {
			return nil, errors.Wrap(err, "failed to build codes, there are incorrect intervals")
		}
		return triads, nil
	}

	tensions, err := f.createTension()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build codes, there are incorrect intervals")
	}
	return tensions, nil
}

//BuildWithName is create tension code with name
//Use this function to generate code whose name is difficult to guess. (e.g. Cm7b5 and Cdim7)
//For triad codes, the name will be ignored.
func (f Factory) BuildWithName(name string) ([]Code, error) {
	if len(f.intervals) == 3 {
		triads, err := f.createTriads()
		if err != nil {
			return nil, errors.Wrap(err, "failed to build codes, there are incorrect intervals")
		}
		return triads, nil
	}

	tensions, err := f.createTensionWithName(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build codes, there are incorrect intervals")
	}
	return tensions, nil
}

func (f Factory) createTriads() ([]Code, error) {
	allNotes := scale.AllNotes()
	triads := make([]Code, len(allNotes))
	for i, root := range scale.AllNotes() {
		triad, err := NewTriad(&root, f.intervals)
		if err != nil {
			return nil, err
		}
		triads[i] = triad
	}

	return triads, nil
}

func (f Factory) createTension() ([]Code, error) {
	allNotes := scale.AllNotes()
	tensions := make([]Code, len(allNotes))
	for i, root := range scale.AllNotes() {
		tension, err := NewTensionCode(&root, f.intervals)
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
		tension, err := NewTensionCodeWithName(root.String() + name, &root, f.intervals)
		if err != nil {
			return nil, err
		}
		tensions[i] = tension
	}

	return tensions, nil
}

func getNotes(root scale.Note, intervals []scale.Interval) ([]scale.Note, error) {
	notes := make([]scale.Note, len(intervals))
	for i, interval := range intervals {
		n, err := root.GetIntervalNote(&interval)
		if err != nil {
			return nil, err
		}
		notes[i] = *n
	}

	return notes, nil
}
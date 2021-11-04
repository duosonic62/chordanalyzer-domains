package chord

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

type Factory struct {
	triad   TriadType
	tension []scale.Interval
}

func NewChordFactory(triad TriadType, tension []scale.Interval) *Factory {
	return &Factory{
		triad:   triad,
		tension: tension,
	}
}

//Build is create chord instances slice
func (f Factory) Build() ([]Chord, error) {
	if len(f.tension) == 0 {
		return nil, errors.New("failed to build chords, there are no intervals")
	}

	tensions, err := f.createTensionChord()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build chords, there are incorrect intervals")
	}
	return tensions, nil
}

//BuildWithName is create tension chord with name
//Use this function to generate chord whose name is difficult to guess. (e.g. Cm7b5 and Cdim7)
//For triad chords, the name will be ignored.
func (f Factory) BuildWithName(name string) ([]Chord, error) {
	if len(f.tension) == 0 {
		return nil, errors.New("failed to build chords, there are no intervals")
	}

	tensions, err := f.createTensionWithName(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build chords, there are incorrect intervals")
	}
	return tensions, nil
}

func (f Factory) createTensionChord() ([]Chord, error) {
	allNotes := scale.AllNotesInOctave()
	tensions := make([]Chord, len(allNotes))
	for i, root := range scale.AllNotesInOctave() {
		tension, err := NewTensionChord(&root, f.triad, f.tension)
		if err != nil {
			return nil, err
		}
		tensions[i] = tension
	}

	return tensions, nil
}

func (f Factory) createTensionWithName(name string) ([]Chord, error) {
	allNotes := scale.AllNotesInOctave()
	tensions := make([]Chord, len(allNotes))
	for i, root := range scale.AllNotesInOctave() {
		tension, err := NewTensionChordWithName(name, &root, f.triad, f.tension)
		if err != nil {
			return nil, err
		}
		tensions[i] = tension
	}

	return tensions, nil
}

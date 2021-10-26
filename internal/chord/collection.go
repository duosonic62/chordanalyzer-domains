package chord

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
)

type Collection interface {
	Get(chordName string) Chord
 	Filter(filter func(chord Chord) bool) Collection
	ForEach(do func(chord Chord))
	ToSlice() []Chord
}

type CollectionFactory struct {
	chords []InOctave
}

func NewCollectionFactory() (*CollectionFactory, error) {
	chords, err := buildTriads()
	if err != nil {
		return nil, err
	}
	return &CollectionFactory{
		chords: chords,
	}, nil
}

func (f CollectionFactory) Append(chordsInOctave *InOctave) *CollectionFactory {
	f.chords = append(f.chords, *chordsInOctave)
	return &f
}

func (f CollectionFactory) Build() Collection {
	return collection{
		allCodes: f.chords,
	}
}

func buildTriads() ([]InOctave, error) {
	triadInOctaves := make([]InOctave, len(AllTriadTypes))
	for i, triadType := range AllTriadTypes {
		triads, err := buildTriad(triadType)
		if err != nil {
			return nil, err
		}

		triadInOctave, err := NewCodesInOctave(triads)
		if err != nil {
			return nil, err
		}

		triadInOctaves[i] = *triadInOctave
	}

	return triadInOctaves, nil
}

func buildTriad(triadType TriadType) ([]Chord, error) {
	triads := make([]Chord, len(scale.AllNotes()))
	for i, root := range scale.AllNotes() {
		triad, err := NewTriadFrom(&root, triadType)
		if err != nil {
			return nil, err
		}

		triads[i] = *triad
	}

	return triads, nil
}

func NewCodesInOctave(chords []Chord) (*InOctave, error) {
	if len(chords) != scale.NoteCount {
		return nil, errors.New("chords in octave must contain " + strconv.Itoa(scale.NoteCount) + " chords")
	}
	return &InOctave{
		Codes: chords,
	}, nil
}

type InOctave struct {
	Codes []Chord
}

type collection struct {
	allCodes []InOctave
}

func (c collection) Get(chordName string) Chord {
	chords := c.Filter(func(chord Chord) bool {
		return chord.Name() == chordName
	})

	var target Chord
	chords.ForEach(func(chord Chord) {
		target = chord
	})

	return target
}

func (c collection) Filter(filter func(chord Chord) bool) Collection {
	var filtered []Chord
	//TODO change to concurrent chord
	for _, inOctave := range c.allCodes {
		for _, chord := range inOctave.Codes {
			if filter(chord) {
				filtered = append(filtered, chord)
			}
		}
	}

	return stream{chords: filtered}
}

func (c collection) ForEach(do func(chord Chord)) {
	//TODO change to concurrent chord
	for _, inOctave := range c.allCodes {
		for _, chord := range inOctave.Codes {
			do(chord)
		}
	}
}

func (c collection) ToSlice() []Chord {
	chords := make([]Chord, len(c.allCodes) * scale.NoteCount)

	count := 0
	c.ForEach(func(chord Chord) {
		chords[count] = chord
		count++
	})

	return chords
}

type stream struct {
	chords []Chord
}

func (c stream) Get(chordName string) Chord {
	chords := c.Filter(func(chord Chord) bool {
		return chord.Name() == chordName
	})

	var target Chord
	chords.ForEach(func(chord Chord) {
		target = chord
	})

	return target
}

func (c stream) Filter(filter func(chord Chord) bool) Collection {
	var filtered []Chord
	//TODO change to concurrent chord
	for _, chord := range c.chords {
		if filter(chord) {
			filtered = append(filtered, chord)
		}
	}

	return stream{chords: filtered}
}

func (c stream) ForEach(do func(chord Chord)) {
	//TODO change to concurrent chord
	for _, chord := range c.chords {
		do(chord)
	}
}

func (c stream) ToSlice() []Chord {
	return c.chords
}

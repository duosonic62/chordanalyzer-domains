package chord

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
)

type Collection interface {
	Get(codeName string) Code
 	Filter(filter func(code Code) bool) Collection
	ForEach(do func(code Code))
	ToSlice() []Code
}

type CollectionFactory struct {
	codes []InOctave
}

func NewCollectionFactory() (*CollectionFactory, error) {
	codes, err := buildTriads()
	if err != nil {
		return nil, err
	}
	return &CollectionFactory{
		codes: codes,
	}, nil
}

func (f CollectionFactory) Append(codesInOctave *InOctave) *CollectionFactory {
	f.codes = append(f.codes, *codesInOctave)
	return &f
}

func (f CollectionFactory) Build() Collection {
	return collection{
		allCodes: f.codes,
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

func buildTriad(triadType TriadType) ([]Code, error) {
	triads := make([]Code, len(scale.AllNotes()))
	for i, root := range scale.AllNotes() {
		triad, err := NewTriadFrom(&root, triadType)
		if err != nil {
			return nil, err
		}

		triads[i] = *triad
	}

	return triads, nil
}

func NewCodesInOctave(codes []Code) (*InOctave, error) {
	if len(codes) != scale.NoteCount {
		return nil, errors.New("codes in octave must contain " + strconv.Itoa(scale.NoteCount) + " codes")
	}
	return &InOctave{
		Codes: codes,
	}, nil
}

type InOctave struct {
	Codes []Code
}

type collection struct {
	allCodes []InOctave
}

func (c collection) Get(codeName string) Code {
	codes := c.Filter(func(code Code) bool {
		return code.Name() == codeName
	})

	var target Code
	codes.ForEach(func(code Code) {
		target = code
	})

	return target
}

func (c collection) Filter(filter func(code Code) bool) Collection {
	var filtered []Code
	//TODO change to concurrent code
	for _, inOctave := range c.allCodes {
		for _, code := range inOctave.Codes {
			if filter(code) {
				filtered = append(filtered, code)
			}
		}
	}

	return stream{codes: filtered}
}

func (c collection) ForEach(do func(code Code)) {
	//TODO change to concurrent code
	for _, inOctave := range c.allCodes {
		for _, code := range inOctave.Codes {
			do(code)
		}
	}
}

func (c collection) ToSlice() []Code {
	codes := make([]Code, len(c.allCodes) * scale.NoteCount)

	count := 0
	c.ForEach(func(code Code) {
		codes[count] = code
		count++
	})

	return codes
}

type stream struct {
	codes []Code
}

func (c stream) Get(codeName string) Code {
	codes := c.Filter(func(code Code) bool {
		return code.Name() == codeName
	})

	var target Code
	codes.ForEach(func(code Code) {
		target = code
	})

	return target
}

func (c stream) Filter(filter func(code Code) bool) Collection {
	var filtered []Code
	//TODO change to concurrent code
	for _, code := range c.codes {
		if filter(code) {
			filtered = append(filtered, code)
		}
	}

	return stream{codes: filtered}
}

func (c stream) ForEach(do func(code Code)) {
	//TODO change to concurrent code
	for _, code := range c.codes {
		do(code)
	}
}

func (c stream) ToSlice() []Code {
	return c.codes
}

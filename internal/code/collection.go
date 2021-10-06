package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
)

type Collection interface {
	Get(codeName string) *Code
 	Filter(filter func(code Code) bool) []Code
	ForEach(do func(code Code))
}

type CollectionFactory struct {
	codes []InOctave
}

func NewCollectionFactory() CollectionFactory {
	return CollectionFactory{
		codes: []InOctave{},
	}
}

func (f CollectionFactory) Append(codesInOctave *InOctave) CollectionFactory {
	f.codes = append(f.codes, *codesInOctave)
	return f
}

func (f CollectionFactory) Build() Collection {
	return collection{
		allCodes: f.codes,
	}
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

func (c collection) Get(codeName string) *Code {
	codes := c.Filter(func(code Code) bool {
		return code.Name() == codeName
	})

	if len(codes) == 0 {
		return nil
	}

	return &codes[0]
}

func (c collection) Filter(filter func(code Code) bool) []Code {
	var filtered []Code
	//TOD change to concurrent code
	for _, inOctave := range c.allCodes {
		for _, code := range inOctave.Codes {
			if filter(code) {
				filtered = append(filtered, code)
			}
		}
	}

	return filtered
}

func (c collection) ForEach(do func(code Code)) {
	//TOD change to concurrent code
	for _, inOctave := range c.allCodes {
		for _, code := range inOctave.Codes {
			do(code)
		}
	}
}

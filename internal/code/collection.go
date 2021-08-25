package code

import (
	"errors"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
)

type Collection interface {
	Filter(filter func(code Code) bool) []Code
}

type CollectionFactory struct {
	codes []InOctave
}

func NewCollectionFactory() CollectionFactory {
	return CollectionFactory{
		codes: []InOctave{},
	}
}

func (f CollectionFactory) Append(codesInOctave InOctave) CollectionFactory {
	f.codes = append(f.codes, codesInOctave)
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

func (c collection) Filter(filter func(code Code) bool) []Code {
	var filtered []Code
	for _, inOctave := range c.allCodes {
		for _, code := range inOctave.Codes {
			if filter(code) {
				filtered = append(filtered, code)
			}
		}
	}

	return filtered
}




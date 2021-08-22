package code

type Collection interface {
	Filter(f func(code Code) bool) []Code
}

type CollectionFactory struct {
	codes []InOctave
}

func NewCollectionFactory() CollectionFactory {
	return CollectionFactory{
		codes: []InOctave{},
	}
}

func (f CollectionFactory) Append(octave InOctave) CollectionFactory {
	f.codes = append(f.codes, octave)
	return f
}

func (f CollectionFactory) Build() Collection {
	return collection{
		allCodes: f.codes,
	}
}

type InOctave struct {
	Name string
	Codes []Code
}

type collection struct {
	allCodes []InOctave
}

func (c collection) Filter(f func(code Code) bool) []Code {
	var filtered []Code
	for _, inOctave := range c.allCodes {
		for _, code := range inOctave.Codes {
			if f(code) {
				filtered = append(filtered, code)
			}
		}
	}

	return filtered
}




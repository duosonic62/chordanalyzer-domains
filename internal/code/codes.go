package code

type Collection interface {
	Filter(f func(code Code) bool) []Code
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




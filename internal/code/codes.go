package code

type Collection interface {
	Filter(f func(code Code) (bool, error)) ([]Code, error)
}

type InOctave struct {
	Name string
	Codes []Code
}

type collection struct {
	allCodes []InOctave
}

func (c collection) Filter(f func(code Code) (bool, error)) ([]Code, error) {
	var filtered []Code
	for _, inOctave := range c.allCodes {
		for _, code := range inOctave.Codes {
			ok, err := f(code)
			if err != nil {
				return nil, err
			}
			if ok {
				filtered = append(filtered, code)
			}
		}
	}

	return filtered, nil
}




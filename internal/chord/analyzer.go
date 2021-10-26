package chord

type Analyzer interface {
	AnalyzeIncludedCodes(target Chord) []Chord
	AllCodes() Collection
}

type analyzer struct {
	collection Collection
}

func NewAnalyzer(collection Collection) Analyzer {
	return analyzer{collection: collection}
}

func (a analyzer) AnalyzeIncludedCodes(target Chord) []Chord {
	return a.collection.Filter(func(code Chord) bool {
		// 自分自身は省く
		return target.Contains(code) && target.Name() != code.Name()
	}).ToSlice()
}

func (a analyzer) AllCodes() Collection {
	return a.collection
}
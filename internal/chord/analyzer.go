package chord

type Analyzer interface {
	AnalyzeIncludedChords(target Chord) []Chord
	AllChords() Collection
}

type analyzer struct {
	collection Collection
}

func NewAnalyzer(collection Collection) Analyzer {
	return analyzer{collection: collection}
}

func (a analyzer) AnalyzeIncludedChords(target Chord) []Chord {
	return a.collection.Filter(func(chord Chord) bool {
		// 自分自身は省く
		return target.Contains(chord) && target.Name() != chord.Name()
	}).ToSlice()
}

func (a analyzer) AllChords() Collection {
	return a.collection
}
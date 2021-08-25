package code

type Analyzer interface {
	AnalyzeIncludedCodes(target Code) []Code
}

type analyzer struct {
	collection Collection
}

func NewAnalyzer(collection Collection) Analyzer {
	return analyzer{collection: collection}
}

func (a analyzer) AnalyzeIncludedCodes(target Code) []Code {
	return a.collection.Filter(func(code Code) bool {
		return target.Contains(code)
	})
}

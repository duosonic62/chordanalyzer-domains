package code

type Analyzer interface {
	AnalyzeIncludedCodes(code *Code) ([]Code, error)
}

type analyzer struct {

}
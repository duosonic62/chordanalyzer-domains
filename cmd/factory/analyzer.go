package factory

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
)

func BuildAnalyzer() (code.Analyzer, error) {
	inputs, err := ReadCodesFromJson("codes.json")
	if err != nil {
		return nil, err
	}

	collectionFactory, err := code.NewCollectionFactory()
	if err != nil {
		return nil, err
	}

	for _, input := range inputs {
		codes, err := CreateCodes(input)
		if err != nil {
			return nil, err
		}

		codesInOctave, err := code.NewCodesInOctave(codes)
		if err != nil {
			return nil, err
		}
		collectionFactory = collectionFactory.Append(codesInOctave)
	}

	collection := collectionFactory.Build()
	analyzer := code.NewAnalyzer(collection)

	return analyzer, nil
}

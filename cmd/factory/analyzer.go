package factory

import (
	"github.com/duosonic62/codanalyzer-domains/internal/chord"
)

func BuildAnalyzer() (chord.Analyzer, error) {
	inputs, err := ReadCodesFromJson("codes.json")
	if err != nil {
		return nil, err
	}

	collectionFactory, err := chord.NewCollectionFactory()
	if err != nil {
		return nil, err
	}

	for _, input := range inputs {
		codes, err := CreateCodes(input)
		if err != nil {
			return nil, err
		}

		codesInOctave, err := chord.NewChordsInOctave(codes)
		if err != nil {
			return nil, err
		}
		collectionFactory = collectionFactory.Append(codesInOctave)
	}

	collection := collectionFactory.Build()
	analyzer := chord.NewAnalyzer(collection)

	return analyzer, nil
}

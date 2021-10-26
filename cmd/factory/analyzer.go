package factory

import (
	"github.com/duosonic62/chordanalyzer-domains/internal/chord"
)

func BuildAnalyzer() (chord.Analyzer, error) {
	inputs, err := ReadChordsFromJson("chords.json")
	if err != nil {
		return nil, err
	}

	collectionFactory, err := chord.NewCollectionFactory()
	if err != nil {
		return nil, err
	}

	for _, input := range inputs {
		chords, err := CreateChords(input)
		if err != nil {
			return nil, err
		}

		chordsInOctave, err := chord.NewChordsInOctave(chords)
		if err != nil {
			return nil, err
		}
		collectionFactory = collectionFactory.Append(chordsInOctave)
	}

	collection := collectionFactory.Build()
	analyzer := chord.NewAnalyzer(collection)

	return analyzer, nil
}

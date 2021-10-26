package factory

import (
	"github.com/duosonic62/codanalyzer-domains/internal/chord"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

func CreateChords(chordInput ChordInput) ([]chord.Chord, error) {
	tension, err := stringsToIntervals(chordInput.Tension)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate chord")
	}
	triadType, err := stringsToTriad(chordInput.Triad)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate chord")
	}

	factory := chord.NewChordFactory(triadType, tension)
	// If the chord name exists, specify it and build it.
	if chordInput.Name != "" {
		chords, err := factory.BuildWithName(chordInput.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate chord")
		}
		return chords, nil
	}

	chords, err := factory.Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate chord")
	}
	return chords, nil
}

func stringsToIntervals(rawIntervals []string) ([]scale.Interval, error) {
	intervals := make([]scale.Interval, len(rawIntervals))
	for i, rawInterval := range rawIntervals {
		interval, err := scale.NewInterval(rawInterval)
		if err != nil {
			return nil, err
		}
		intervals[i] = *interval
	}
	return intervals, nil
}

func stringsToTriad(rawTriad string) (chord.TriadType, error) {
	switch rawTriad {
	case "Major":
		return chord.Major, nil
	case "Minor":
		return chord.Minor, nil
	case "Augment":
		return chord.Augment, nil
	case "Diminish":
		return chord.Diminish, nil
	case "MajorB5":
		return chord.MajorB5, nil
	case "Sus2":
		return chord.Sus2, nil
	case "Sus4":
		return chord.Sus4, nil
	}

	return "", errors.New(rawTriad + " is not a valid triad" )
}

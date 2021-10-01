package factory

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/pkg/errors"
)

func CreateCodes(codeInput CodeInput) ([]code.Code, error) {
	tension, err := stringsToIntervals(codeInput.Tension)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate code")
	}
	triadType, err := stringsToTriad(codeInput.Triad)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate code")
	}

	factory := code.NewCodeFactory(triadType, tension)
	// If the code name exists, specify it and build it.
	if codeInput.Name != "" {
		codes, err := factory.BuildWithName(codeInput.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate code")
		}
		return codes, nil
	}

	codes, err := factory.Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate code")
	}
	return codes, nil
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

func stringsToTriad(rawTriad string) (code.TriadType, error) {
	switch rawTriad {
	case string(code.Major):
		return code.Major, nil
	case string(code.Minor):
		return code.Minor, nil
	case string(code.Augment):
		return code.Augment, nil
	case string(code.Diminish):
		return code.Diminish, nil
	case string(code.MajorB5):
		return code.MajorB5, nil
	case string(code.Sus2):
		return code.Sus2, nil
	case string(code.Sus4):
		return code.Sus4, nil
	}

	return "", errors.New("")
}

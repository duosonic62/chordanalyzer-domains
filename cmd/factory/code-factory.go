package factory

import (
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
)

func CreateCodes(codeInput CodeInput) ([]code.Code, error) {
	intervals, err := stringsToIntervals(codeInput.Intervals)
	if err != nil {
		return nil, err
	}

	factory :=  code.NewFactory(intervals)
	// If the code name exists, specify it and build it.
	if codeInput.Name != "" {
		return factory.BuildWithName(codeInput.Name)
	}

	return factory.Build()
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

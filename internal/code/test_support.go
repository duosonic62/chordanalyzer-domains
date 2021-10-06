package code

import "github.com/duosonic62/codanalyzer-domains/internal/scale"

var majorFactory = NewFactory([]scale.Interval{
	scale.Intervals.R,
	scale.Intervals.Major3,
	scale.Intervals.Perfect5,
})

var majorCodes, _ = majorFactory.Build()
var major, _ = NewCodesInOctave(majorCodes)
//var majorCollection = NewCollectionFactory().Append(major).Build()

var minorvarFactory = NewFactory([]scale.Interval{
	scale.Intervals.R,
	scale.Intervals.Minor3,
	scale.Intervals.Perfect5,
})

var minorCodes, _ = minorvarFactory.Build()
var minor, _ = NewCodesInOctave(minorCodes)
//var minorCollection = NewCollectionFactory().Append(minor).Build()

var major7Factory = NewFactory([]scale.Interval{
	scale.Intervals.R,
	scale.Intervals.Major3,
	scale.Intervals.Perfect5,
	scale.Intervals.Major7,
})

var major7codes, _ = major7Factory.Build()
var major7, _ = NewCodesInOctave(major7codes)
var factory, _ = NewCollectionFactory()
var major7Collection = factory.Append(major7).Build()

var collections = factory.Append(major).Append(minor).Append(major7).Build()
package code

import "github.com/duosonic62/codanalyzer-domains/internal/scale"

var major7Factory = NewCodeFactory(Major, []scale.Interval{scale.Intervals.Major7})

var major7codes, _ = major7Factory.Build()
var major7, _ = NewCodesInOctave(major7codes)
var factory, _ = NewCollectionFactory()
var major7Collection = factory.Append(major7).Build()

var collections = major7Collection

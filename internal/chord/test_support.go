package chord

import "github.com/duosonic62/codanalyzer-domains/internal/scale"

var major7Factory = NewChordFactory(Major, []scale.Interval{scale.Intervals.Major7})

var major7chords, _ = major7Factory.Build()
var major7, _ = NewChordsInOctave(major7chords)
var factory, _ = NewCollectionFactory()
var major7Collection = factory.Append(major7).Build()

var collections = major7Collection

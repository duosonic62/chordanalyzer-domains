package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
	"strings"
	"testing"
)

var factory = NewFactory([]scale.Interval{
	scale.Intervals.R,
	scale.Intervals.Major3,
	scale.Intervals.Perfect5,
	scale.Intervals.Major7,
})

var codes, _ = factory.Build()
var major7, _ = NewCodesInOctave(codes)
var major7Collection = NewCollectionFactory().Append(major7).Build()

func TestCollection_Filter(t *testing.T) {
	actual := major7Collection.Filter(func(code Code) bool {
		return strings.Contains(code.Name(), "C")
	})

	if len(actual) !=2 {
		t.Error("Expected: CM7 and C#M7 but not")
	}
}

func TestCollection_ForEach(t *testing.T) {
	count := 0
	major7Collection.ForEach(func(code Code) {
		count++
	})

	if count != 12 {
		t.Error("Expected: 12 actions but actual: " + strconv.Itoa(12))
	}
}
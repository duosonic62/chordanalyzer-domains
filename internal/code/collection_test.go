package code

import (
	"strconv"
	"strings"
	"testing"
)

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
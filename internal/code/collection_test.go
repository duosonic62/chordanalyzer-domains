package code

import (
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"strconv"
	"strings"
	"testing"
)

func TestCollection_Filter(t *testing.T) {
	actual := major7Collection.Filter(func(code Code) bool {
		return strings.Contains(code.Name(), "C")
	})

	if len(actual) !=16 {
		t.Error("Expected: C* and C#* but not")
	}
}

func TestCollection_ForEach(t *testing.T) {
	count := 0
	major7Collection.ForEach(func(code Code) {
		count++
	})

	if count /len(scale.AllNotes()) != len(AllTriadTypes) + 1 {
		t.Error("Expected: " + strconv.Itoa(len(AllTriadTypes) + 1) + " actions but actual: " + strconv.Itoa(count /len(scale.AllNotes())))
	}
}
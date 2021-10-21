package cli

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/code"
)

func PrintCode(code code.Code)  {
	fmt.Println(code.Name())
	fmt.Print("Notes     | ")
	for _, note := range code.Notes() {
		fmt.Printf("%5s |", note.String())
	}
	fmt.Println()
	fmt.Print("Intervals | ")
	for _, interval := range code.Intervals() {
		fmt.Printf("%5s |", interval.String())
	}
	fmt.Println()
}

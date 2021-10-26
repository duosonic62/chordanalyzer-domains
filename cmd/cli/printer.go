package cli

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/chord"
)

func PrintChord(chord chord.Chord)  {
	fmt.Println(chord.Name())
	fmt.Print("Notes     | ")
	for _, note := range chord.Notes() {
		fmt.Printf("%5s |", note.String())
	}
	fmt.Println()
	fmt.Print("Intervals | ")
	for _, interval := range chord.Intervals() {
		fmt.Printf("%5s |", interval.String())
	}
	fmt.Println()
}

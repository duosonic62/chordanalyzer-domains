package cli

import (
	"fmt"
	"github.com/duosonic62/chordanalyzer-domains/internal/chord"
	"github.com/duosonic62/chordanalyzer-domains/internal/scale"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use: "list",
		Short: "List all available chords\n Show only chords whose root note is C",
		Run: runList,
	}
)

func runList(cmd *cobra.Command, args []string) {
	listAction()
}

func listAction() {
	Analyzer.AllChords().Filter(func(chord chord.Chord) bool {
		return chord.Root().Equals(&scale.Notes.C)
	}).ForEach(func(chord chord.Chord) {
		fmt.Println(chord.Name())
	})
}

func init()  {
	RootCmd.AddCommand(listCmd)
}
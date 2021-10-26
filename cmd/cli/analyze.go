package cli

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	analyzeCmd = &cobra.Command{
		Use: "analyze",
		Short: "Analyze chord",
		Run: runAnalyze,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("Require a valid chord.")
			}
			return nil
		},
	}
)

func runAnalyze(cmd *cobra.Command, args []string) {
	analyzeAction(args[0])
}

func analyzeAction(chord string) {
	target := Analyzer.AllChords().Get(chord)
	if target == nil {
		Exit(errors.New(chord+" is a invalid chord\n Please check valid chords with list command"), 2)
		return
	}

	fmt.Println("Analyzed Chord")
	PrintChord(target)

	fmt.Println(target.Name() + " include...")
	for _, chord := range Analyzer.AnalyzeIncludedChords(target) {
		PrintChord(chord)
	}
}

func init() {
	RootCmd.AddCommand(analyzeCmd)
}

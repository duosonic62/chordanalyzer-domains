package cli

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	analyzeCmd = &cobra.Command{
		Use: "Analyze chord",
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
	target := Analyzer.AllCodes().Get(chord)
	if target == nil {
		Exit(errors.New(chord+" is a invalid chord\n Please check valid chords with list command"), 2)
	}

	fmt.Println("Analyzed Code")
	fmt.Println(target.Info())

	fmt.Println(target.Name() + " include...")
	for _, code := range Analyzer.AnalyzeIncludedCodes(target) {
		fmt.Println(code.Info())
	}
}

func init() {
	RootCmd.AddCommand(analyzeCmd)
}

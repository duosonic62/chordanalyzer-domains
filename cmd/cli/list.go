package cli

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/chord"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use: "list",
		Short: "List all available codes\n Show only chords whose root note is C",
		Run: runList,
	}
)

func runList(cmd *cobra.Command, args []string) {
	listAction()
}

func listAction() {
	Analyzer.AllCodes().Filter(func(code chord.Code) bool {
		return code.Root().Equals(&scale.Notes.C)
	}).ForEach(func(code chord.Code) {
		fmt.Println(code.Name())
	})
}

func init()  {
	RootCmd.AddCommand(listCmd)
}
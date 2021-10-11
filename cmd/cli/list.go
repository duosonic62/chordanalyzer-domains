package cli

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use: "list all available codes",
		Run: runList,
	}
)

func runList(cmd *cobra.Command, args []string) {
	listAction()
}

func listAction() {
	Analyzer.AllCodes().ForEach(func(code code.Code) {
		fmt.Println(code.Name())
	})
}

func init()  {
	RootCmd.AddCommand(listCmd)
}
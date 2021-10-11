package cli

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/spf13/cobra"
	"os"
)

var (
	// RootCmd defines root command
	RootCmd = &cobra.Command{
		Use: "code-analyze",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	Analyzer code.Analyzer
)

func Run()  {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
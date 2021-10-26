package cli

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/internal/chord"
	"github.com/spf13/cobra"
	"os"
)

var (
	// RootCmd defines root command
	RootCmd = &cobra.Command{
		Use: "chord-analyze",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Usage()
			if err != nil {
				Exit(err, 2)
			}
		},
	}
	Analyzer chord.Analyzer
)

func Run()  {
	err := RootCmd.Execute()
	if err != nil {
		Exit(err, 2)
	}
}

func Exit(err error, exit int)  {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(exit)
}
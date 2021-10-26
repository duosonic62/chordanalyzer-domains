package main

import (
	"fmt"
	"github.com/duosonic62/chordanalyzer-domains/cmd/cli"
	"github.com/duosonic62/chordanalyzer-domains/cmd/factory"
	"os"
)

func main() {
	analyzer, err := factory.BuildAnalyzer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cli.Analyzer = analyzer
	cli.Run()
}

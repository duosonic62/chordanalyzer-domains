package main

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/cmd/factory"
)

func main() {
	inputs, err := factory.ReadCodesFromJson("codes.json")
	if err != nil {
		panic(err)
	}

	for _, input := range inputs {
		codes, err := factory.CreateCodes(input)
		if err != nil {
			panic(err)
		}

		fmt.Println(codes[0].Name())
	}
}

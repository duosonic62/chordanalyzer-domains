package main

import (
	"fmt"
	"github.com/duosonic62/codanalyzer-domains/cmd/factory"
	"github.com/duosonic62/codanalyzer-domains/internal/code"
	"github.com/duosonic62/codanalyzer-domains/internal/scale"
)

func main() {
	inputs, err := factory.ReadCodesFromJson("codes.json")
	if err != nil {
		panic(err)
	}

	collectionFactory, err := code.NewCollectionFactory()
	if err != nil {
		panic(err)
	}

	for _, input := range inputs {
		codes, err := factory.CreateCodes(input)
		if err != nil {
			panic(err)
		}

		codesInOctave, err := code.NewCodesInOctave(codes)
		if err != nil {
			panic(err)
		}
		collectionFactory = collectionFactory.Append(codesInOctave)
	}

	collection := collectionFactory.Build()
	collection.ForEach(func(code code.Code) {
		if code.Root().Equals(&scale.Notes.C) {
			fmt.Println(code.Name())
		}
	})
	analyzer := code.NewAnalyzer(collection)

	target := collection.Get("Cm7b5")
	if target != nil {
		analyzedCodes := analyzer.AnalyzeIncludedCodes(*target)
		for _, analyzedCode := range analyzedCodes {
			fmt.Println(analyzedCode)
		}
	}
}

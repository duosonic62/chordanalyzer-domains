package factory

import (
	"encoding/json"
	"io/ioutil"
)

type CodeInputCollection struct {
	Version string      `json:"version"`
	Codes   []CodeInput `json:"codes"`
}

type CodeInput struct {
	Name      string   `json:"name"`
	Intervals []string `json:"intervals"`
}

func ReadCodesFromJson(filePath string) ([]CodeInput, error) {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var codeInputCollection CodeInputCollection
	err = json.Unmarshal(raw, &codeInputCollection)
	if err != nil {
		return nil, err
	}

	return codeInputCollection.Codes, nil
}

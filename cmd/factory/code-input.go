package factory

import (
	"encoding/json"
	"io/ioutil"
)

type ChordInputCollection struct {
	Version string      `json:"version"`
	Chords   []ChordInput `json:"chords"`
}

type ChordInput struct {
	Name    string   `json:"name"`
	Triad   string   `json:"triad"`
	Tension []string `json:"tension"`
}

func ReadChordsFromJson(filePath string) ([]ChordInput, error) {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var chordInputCollection ChordInputCollection
	err = json.Unmarshal(raw, &chordInputCollection)
	if err != nil {
		return nil, err
	}

	return chordInputCollection.Chords, nil
}

package scale

import "errors"

type Code struct {
	Name  string
	Root  *Note
	Notes *[]Note
}

func NewCode(notes []Note) (*Code, error) {
	if len(notes) < 3 {
		return nil, errors.New("need at least 3 chord tones")
	}

	// ルートノート
	root := notes[0]

	// トライアドを検出する
	traid := notes[0:2]

	// トライアド以外を検出する
	tention := notes[3:]

	return nil, nil
}
package scale

type Triad struct {
	name  string
	root  *Note
	notes []Note
}

func (t Triad) Name() string {
	return t.name
}

func (t Triad) Root() *Note {
	return t.root
}

func (t Triad) Notes() []Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]Note, len(t.notes), cap(t.notes))
	copy(notes, t.notes)
	return notes
}


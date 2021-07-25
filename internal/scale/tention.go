package scale

type Tension struct {
	name  string
	triad *Triad
	tensionNotes []Note
}

func (t Tension) Name() string {
	return t.name
}

func (t Tension) Root() *Note {
	return t.triad.root
}

func (t Tension) Notes() []Note {
	// 内容を書き換えられても良いようにスライスのコピーを渡す
	notes := make([]Note, t.notesNum())
	copy(notes, append(t.triad.Notes(), t.tensionNotes...))
	return notes
}

func (t Tension) notesNum() int {
	return len(t.tensionNotes) + len(t.triad.notes)
}
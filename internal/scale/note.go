package scale

import "errors"

type Note struct {
	name          string
	intervalFromC int
}

// クロマチックスケールの音数
var NoteCount = 12

var Notes = struct {
	C  Note
	CS Note
	Db Note
	D  Note
	DS Note
	Eb Note
	E  Note
	F  Note
	FS Note
	Gb Note
	G  Note
	GS Note
	Ab Note
	A  Note
	AS Note
	Bb Note
	B  Note
}{
	C: Note{
		name:          "C",
		intervalFromC: 0,
	},
	CS: Note{
		name:          "C#",
		intervalFromC: 1,
	},
	Db: Note{
		name:          "Db",
		intervalFromC: 1,
	},
	D: Note{
		name:          "D",
		intervalFromC: 2,
	},
	DS: Note{
		name:          "D#",
		intervalFromC: 3,
	},
	Eb: Note{
		name:          "Eb",
		intervalFromC: 3,
	},
	E: Note{
		name:          "E",
		intervalFromC: 4,
	},
	F: Note{
		name:          "F",
		intervalFromC: 5,
	},
	FS: Note{
		name:          "F#",
		intervalFromC: 6,
	},
	Gb: Note{
		name:          "Gb",
		intervalFromC: 6,
	},
	G: Note{
		name:          "G",
		intervalFromC: 7,
	},
	GS: Note{
		name:          "G#",
		intervalFromC: 8,
	},
	Ab: Note{
		name:          "Ab",
		intervalFromC: 8,
	},
	A: Note{
		name:          "A",
		intervalFromC: 9,
	},
	AS: Note{
		name:          "A#",
		intervalFromC: 10,
	},
	Bb: Note{
		name:          "Bb",
		intervalFromC: 10,
	},
	B: Note{
		name:          "B",
		intervalFromC: 11,
	},
}

//CalculateInterval は二音間のインターバルを計算する
func (n Note) CalculateInterval(compareNNote Note) (*Interval, error) {
	if compareNNote.intervalFromC >= n.intervalFromC {
		return intervalFromNumber(compareNNote.intervalFromC - n.intervalFromC)
	}

	return intervalFromNumber(compareNNote.intervalFromC + (NoteCount - n.intervalFromC))
}

func (n Note) CalculateTensionInterval(compareNNote Note) (*Interval, error) {
	if compareNNote.intervalFromC >= n.intervalFromC {
		return tensionFromNumber(compareNNote.intervalFromC - n.intervalFromC)
	}

	return tensionFromNumber(compareNNote.intervalFromC + (NoteCount - n.intervalFromC))
}

//GetIntervalNote はInterval分離れたノートを取得する
func (n Note) GetIntervalNote(interval *Interval) (*Note, error) {
	noteInterval := n.intervalFromC + interval.value
	if noteInterval >= NoteCount {
		noteInterval = noteInterval - NoteCount
	}

	note, err := noteFromInterval(noteInterval)
	if err != nil {
		return nil, err
	}

	return note, nil
}

//Equals is strict compare notes
func (n Note) Equals(other *Note) bool {
	if other == nil {
		return false
	}

	// 名前と音階が一緒なら同値
	return n.name == other.name && n.intervalFromC == other.intervalFromC
}

//Equivalent is compare notes (not case sensitive, A# and Bb)
func (n Note) Equivalent(other *Note) bool {
	if other == nil {
		return false
	}

	return n.intervalFromC == other.intervalFromC
}

func (n Note) String() string {
	return n.name
}

func noteFromInterval(interval int) (*Note, error) {
	switch interval {
	case 0:
		return &Notes.C, nil
	case 1:
		return &Notes.CS, nil
	case 2:
		return &Notes.D, nil
	case 3:
		return &Notes.DS, nil
	case 4:
		return &Notes.E, nil
	case 5:
		return &Notes.F, nil
	case 6:
		return &Notes.FS, nil
	case 7:
		return &Notes.G, nil
	case 8:
		return &Notes.GS, nil
	case 9:
		return &Notes.A, nil
	case 10:
		return &Notes.AS, nil
	case 11:
		return &Notes.B, nil
	default:
		return nil, errors.New("note interval must be 0 - 11")
	}
}

//FromString is to get note from name.
func FromString(noteName string) (*Note, error) {
	for _, note := range AllNotes() {
		if note.String() == noteName {
			return &note, nil
		}
	}
	return nil, errors.New("note doesn't contain " + noteName)
}

func AllNotesInOctave() []Note {
	return [] Note{
		Notes.C,
		Notes.CS,
		Notes.D,
		Notes.DS,
		Notes.E,
		Notes.F,
		Notes.FS,
		Notes.G,
		Notes.GS,
		Notes.A,
		Notes.AS,
		Notes.B,
	}
}

func AllNotes() []Note {
	return [] Note{
		Notes.C,
		Notes.CS,
		Notes.Db,
		Notes.D,
		Notes.DS,
		Notes.Eb,
		Notes.E,
		Notes.F,
		Notes.FS,
		Notes.Gb,
		Notes.G,
		Notes.GS,
		Notes.Ab,
		Notes.A,
		Notes.AS,
		Notes.Bb,
		Notes.B,
	}
}
package scale

import "errors"

type Note struct {
	name          string
	intervalFromC int
}

// クロマチックスケールの音数
var scaleNoteCount = 12

var Notes = struct {
	C  Note
	CS Note
	D  Note
	DS Note
	E  Note
	F  Note
	FS Note
	G  Note
	GS Note
	A  Note
	AS Note
	B  Note
}{
	C: Note{
		name:    "C",
		intervalFromC: 0,
	},
	CS: Note{
		name:    "C#",
		intervalFromC: 1,
	},
	D: Note{
		name:    "D",
		intervalFromC: 2,
	},
	DS: Note{
		name:    "D#",
		intervalFromC: 3,
	},
	E: Note{
		name:    "E",
		intervalFromC: 4,
	},
	F: Note{
		name:    "F",
		intervalFromC: 5,
	},
	FS: Note{
		name:    "F#",
		intervalFromC: 6,
	},
	G: Note{
		name:    "G",
		intervalFromC: 7,
	},
	GS: Note{
		name:    "G#",
		intervalFromC: 8,
	},
	A: Note{
		name:    "A",
		intervalFromC: 9,
	},
	AS: Note{
		name:    "A#",
		intervalFromC: 10,
	},
	B: Note{
		name:    "B",
		intervalFromC: 11,
	},
}

//CalculateInterval は二音間のインターバルを計算する
func (n Note) CalculateInterval(compareNNote Note) (*Interval, error) {
	if compareNNote.intervalFromC >= n.intervalFromC {
		return intervalFromNumber(compareNNote.intervalFromC - n.intervalFromC)
	}

	return intervalFromNumber(compareNNote.intervalFromC + (scaleNoteCount - n.intervalFromC))
}

//GetIntervalNote はInterval分離れたノートを取得する
func (n Note) GetIntervalNote(interval *Interval) (*Note, error) {
	noteInterval := n.intervalFromC + interval.value
	if noteInterval >= scaleNoteCount {
		noteInterval = noteInterval - scaleNoteCount
	}

	note, err := noteFromInterval(noteInterval)
	if err != nil {
		return nil, err
	}

	return note, nil
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
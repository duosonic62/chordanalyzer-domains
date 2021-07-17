package scale

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
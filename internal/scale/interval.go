package scale

import "errors"

type Interval struct {
	value int
	name  string
	upper bool
}

var Intervals = struct {
	R         Interval
	Minor2    Interval
	Major2    Interval
	Minor3    Interval
	Major3    Interval
	Perfect4  Interval
	Sharp4    Interval
	Flat5     Interval
	Perfect5  Interval
	Sharp5    Interval
	Minor6    Interval
	Major6    Interval
	Sharp6    Interval
	Minor7    Interval
	Major7    Interval
	Flat9     Interval
	Natural9  Interval
	Sharp9    Interval
	Natural11 Interval
	Sharp11   Interval
	Flat13    Interval
	Natural13 Interval
}{
	R: Interval{
		value: 0,
		name:  "R",
		upper: false,
	},
	Minor2: Interval{
		value: 1,
		name:  "m2",
		upper: false,
	},
	Major2: Interval{
		value: 2,
		name:  "M2",
		upper: false,
	},
	Minor3: Interval{
		value: 3,
		name:  "m3",
		upper: false,
	},
	Major3: Interval{
		value: 4,
		name:  "M3",
		upper: false,
	},
	Perfect4: Interval{
		value: 5,
		name:  "P4",
		upper: false,
	},
	Sharp4: Interval{
		value: 6,
		name:  "#4",
		upper: false,
	},
	Flat5: Interval{
		value: 6,
		name:  "b5",
		upper: false,
	},
	Perfect5: Interval{
		value: 7,
		name:  "P5",
		upper: false,
	},
	Sharp5: Interval{
		value: 8,
		name:  "#5",
		upper: false,
	},
	Minor6: Interval{
		value: 8,
		name:  "m6",
		upper: false,
	},
	Major6: Interval{
		value: 9,
		name:  "M6",
		upper: false,
	},
	Sharp6: Interval{
		value: 10,
		name:  "#6",
		upper: false,
	},
	Minor7: Interval{
		value: 10,
		name:  "m7",
		upper: false,
	},
	Major7: Interval{
		value: 11,
		name:  "M7",
		upper: false,
	},
	Flat9: Interval{
		value: 1,
		name:  "b9",
		upper: true,
	},
	Natural9: Interval{
		value: 2,
		name:  "9",
		upper: true,
	},
	Sharp9: Interval{
		value: 2,
		name:  "9",
		upper: true,
	},
	Natural11: Interval{
		value: 5,
		name:  "11",
		upper: true,
	},
	Sharp11: Interval{
		value: 6,
		name:  "#11",
		upper: true,
	},
	Flat13: Interval{
		value: 8,
		name:  "b13",
		upper: true,
	},
	Natural13: Interval{
		value: 9,
		name:  "13",
		upper: true,
	},
}

//IsEquivalent はインターバルが等価であるか比較する
//同オクターブ上でなくても同じ音解であれば等価とする
func (i Interval) IsEquivalent(target *Interval) bool {
	return i.value == target.value
}

//IsEquals はインターバルが等値であるか比較する
//同音でないと等値にならない
func (i Interval) IsEquals(target *Interval) bool {
	return i.value == target.value && i.upper == target.upper
}

//intervalFromNumber は数字からインターバルを復元する
func intervalFromNumber(num int) (*Interval, error) {
	switch num {
	case 0:
		return &Intervals.R, nil
	case 1:
		return &Intervals.Minor2, nil
	case 2:
		return &Intervals.Major2, nil
	case 3:
		return &Intervals.Minor3, nil
	case 4:
		return &Intervals.Major3, nil
	case 5:
		return &Intervals.Perfect4, nil
	case 6:
		return &Intervals.Sharp4, nil
	case 7:
		return &Intervals.Perfect5, nil
	case 8:
		return &Intervals.Sharp5, nil
	case 9:
		return &Intervals.Major6, nil
	case 10:
		return &Intervals.Minor7, nil
	case 11:
		return &Intervals.Major7, nil
	default:
		return nil, errors.New("scale tone interval must be 0 - 11")
	}
}
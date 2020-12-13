package axiom

import (
	"golang_music/internal/util"
)

// Tonal :
func Tonal() *util.Dict {
	return &util.Dict{
		1: "major",
		2: "minor",
	}
}

// TwelveNotes :
func TwelveNotes(tonal uint) *util.Dict {

	if isMinor(tonal) == true {
		return minorNotes()
	}
	return majorNotes()

}

func isMinor(tonal uint) bool {
	t := Tonal()
	r, e := t.ValueOf(tonal)
	if e != nil || r != "minor" {
		return false
	}
	return true
}

func majorNotes() *util.Dict {
	return &util.Dict{
		1:  "C",
		2:  "D-",
		3:  "D",
		4:  "E-",
		5:  "E",
		6:  "F",
		7:  "G-",
		8:  "G",
		9:  "A-",
		10: "A",
		11: "B-",
		12: "B",
	}
}

func minorNotes() *util.Dict {
	return &util.Dict{
		1:  "C",
		2:  "C+",
		3:  "D",
		4:  "D+",
		5:  "E",
		6:  "F",
		7:  "F+",
		8:  "G",
		9:  "G+",
		10: "A",
		11: "A+",
		12: "B",
	}
}

// Degrees :
func Degrees(key uint) *util.Dict {

	d := &util.Dict{
		1:  "1.0",
		2:  "1.5",
		3:  "2.0",
		4:  "2.5",
		5:  "3.0",
		6:  "4.0",
		7:  "4.5",
		8:  "5.0",
		9:  "5.5",
		10: "6.0",
		11: "6.5",
		12: "7.0",
	}

	r := util.Dict{}
	var k uint
	for i := 1; i < 13; i++ {
		k = keyCenterIndex(key, uint(i))
		v, _ := d.ValueOf(uint(i))
		r[k] = v
	}

	return &r
}

func keyCenterIndex(key uint, index uint) uint {

	var d int
	d = int(index)+int(key)-1
	
	if d > 12 {
		return uint(d-12)
	}

	return uint(d)
}

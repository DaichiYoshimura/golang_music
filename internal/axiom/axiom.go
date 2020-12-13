package axiom

import "golang_music/internal/util"

// Tonal :
func Tonal() util.Dict {
	return util.Dict{
		1: "major",
		2: "minor",
	}
}

// TwelveNotes :
func TwelveNotes(tonal uint) util.Dict {

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

func majorNotes() util.Dict {
	return util.Dict{
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

func minorNotes() util.Dict {
	return util.Dict{
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

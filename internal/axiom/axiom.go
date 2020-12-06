package axiom

import "golang_music/internal/util/dict"

// Tonal :
func Tonal() dict.Dict {
	return dict.Dict{
		1: "major",
		2: "minor",
	}
}

// TwelveNotes :
func TwelveNotes(tonal uint) dict.Dict {

	if isMinor(tonal) {
		return minorNotes()
	}
	return majorNotes()

}

func isMinor(tonal uint) bool {
	t := Tonal()
	return t.IndexOf("minor") == tonal
}

func majorNotes() dict.Dict {
	return dict.Dict{
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

func minorNotes() dict.Dict {
	return dict.Dict{
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

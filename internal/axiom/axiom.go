package axiom

func Tonal() map[uint]string {
	return  map[uint]string{
		0:"Major",
		1:"minor",
	}
}

func TwelveNotes(tonal uint,enharmonic bool) map[uint]string {
	return map[uint]string{
		0:"C",
		1:"D-",
		2:"D",
		3:"E-",
		4:"E",
		5:"F",
		6:"G-",
		7:"G",
		8:"A-",
		9:"A",
		10:"B-",
		11:"B",
	}
}


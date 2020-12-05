package axiom

func TonalMap() map[uint]string {
	return  map[uint]string{
		1:"major",
		2:"minor",
	}
}

func TwelveNotesMap(tonal uint) map[uint]string {
	
	major :=map[uint]string{
		1:"C",
		2:"D-",
		3:"D",
		4:"E-",
		5:"E",
		6:"F",
		7:"G-",
		8:"G",
		9:"A-",
		10:"A",
		11:"B-",
		12:"B",
	}

	minor :=map[uint]string{
		1:"C",
		2:"C+",
		3:"D",
		4:"D+",
		5:"E",
		6:"F",
		7:"F+",
		8:"G",
		9:"G+",
		10:"A",
		11:"A+",
		12:"B",
	}

	if TonalMap()[tonal] == TonalMap()[1] {
		return major
	}

	if TonalMap()[tonal] == TonalMap()[2] {
		return minor
	}

	return major
	
}


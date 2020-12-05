package note

type Note struct {
	index uint
	sign string
	degree uint
	octave uint
	duration float32
	tonal uint
	key uint
}

func (note *Note) Sign() string {
	return note.sign
}

func (note *Note) SetSign(sign string){
	note.sign = sign
}

func (note *Note) Degree() uint {
	return note.degree
}

func (note *Note) SetDegree(degree uint){
	note.degree = degree
}

func (note *Note) Octave() uint {
	return note.octave
}

func (note *Note) SetOctave(octave uint){
	note.octave = octave
}

func (note *Note) Duration() float32 {
	return note.duration
}

func (note *Note) SetDuration(duration float32){
	note.duration = duration
}





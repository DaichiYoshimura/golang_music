package note

import "golang_music/internal/tonality"

type Note struct {
	index uint
	sign string
	degree float32
	octave uint
	duration float32
	tonality tonality.Tonality
}

func NewNote(tonality tonality.Tonality) *Note {
	n :=new(Note)
	n.tonality = tonality
	return n
}

func (note *Note) Index() uint {
	return note.index
}

func (note *Note) SetIndex(index uint){
	note.index = index
}

func (note *Note) Sign() string {
	return note.sign
}

func (note *Note) SetSign(sign string){
	note.sign = sign
}

func (note *Note) Degree() float32 {
	return note.degree
}

func (note *Note) SetDegree(degree float32){
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





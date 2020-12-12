package note

import (
	"golang_music/internal/axiom"
	"golang_music/internal/tonality"
)

// Note :
type Note struct {
	index    uint
	symbol   string
	degree   float32
	octave   uint
	duration float32
	tonality tonality.Tonality
}

//New :
func New(tonality *tonality.Tonality) *Note {
	i := new(Note)
	i.tonality = *tonality
	return i
}

// Index :
func (n *Note) Index() uint {
	return n.index
}

// SetIndex :
func (n *Note) SetIndex(index uint) {
	n.index = index
}

// Symbol :
func (n *Note) Symbol() string {
	return n.symbol
}

// SetSymbol :
func (n *Note) SetSymbol(symbol string) {
	n.symbol = symbol
}

// Degree :
func (n *Note) Degree() float32 {
	return n.degree
}

// SetDegree :
func (n *Note) SetDegree(degree float32) {
	n.degree = degree
}

// Octave :
func (n *Note) Octave() uint {
	return n.octave
}

// SetOctave :
func (n *Note) SetOctave(octave uint) {
	n.octave = octave
}

// Duration :
func (n *Note) Duration() float32 {
	return n.duration
}

// SetDuration :
func (n *Note) SetDuration(duration float32) {
	n.duration = duration
}

func (n *Note) tonalIndex() uint {
	return n.tonality.TonalIndex()
}

func (n *Note) setTonalIndex(index uint) {
	n.tonality.SetTonalIndex(index)
}

func (n *Note) keyIndex() uint {
	return n.tonality.KeyIndex()
}

func (n *Note) setKeyIndex(index uint) {
	n.tonality.SetKeyIndex(index)
}

func (n *Note) symbolOfIndex(index uint) string {
	cnv := axiom.TwelveNotes(n.tonalIndex())
	return cnv.ValueOf(index)
}

func (n *Note) symbolOfDegree(degree float32) {

}

func (n *Note) indexOfSymbol(symbol string) uint {
	cnv := axiom.TwelveNotes(n.tonalIndex())
	return cnv.IndexOf(symbol)
}

func (n *Note) indexOfDegree(degree float32) {

}

func (n *Note) degreeOfIndex(index uint) {

}

func (n *Note) degreeOfSymbol(symbol string) {

}

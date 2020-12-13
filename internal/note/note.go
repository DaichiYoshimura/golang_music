package note

import (
	"golang_music/internal/axiom"
	"golang_music/internal/tonality"
	"golang_music/internal/util"
)

// Note :
type Note struct {
	index    uint
	symbol   *util.Item
	degree   *util.Item
	octave   uint
	duration float64
	tonality *tonality.Tonality
}

//New :
func New(t *tonality.Tonality) *Note {
	i := new(Note)
	i.tonality = t
	i.symbol = util.New(axiom.TwelveNotes(i.tonalIndex()))
	i.degree = util.New(axiom.Degrees(i.keyIndex()))
	return i
}

// Index :
func (n *Note) Index() uint {
	return n.index
}

// SetIndex :
func (n *Note) SetIndex(index uint) {
	n.index = index
	n.degree.SetIndex(index)
	n.symbol.SetIndex(index)
}

// Symbol :
func (n *Note) Symbol() string {
	n.reinitialize()
	return n.symbol.Value()
}

// SetSymbol :
func (n *Note) SetSymbol(symbol string) {
	n.symbol.SetValue(symbol)
	n.SetIndex(n.symbol.Index())
	n.degree.SetIndex(n.Index())
}

// Degree :
func (n *Note) Degree() string {
	n.reinitialize()
	return n.degree.Value()
}

// SetDegree :
func (n *Note) SetDegree(degree string) {
	n.degree.SetValue(degree)
	n.SetIndex(n.degree.Index())
	n.symbol.SetIndex(n.Index())
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
func (n *Note) Duration() float64 {
	return n.duration
}

// SetDuration :
func (n *Note) SetDuration(duration float64) {
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

func (n *Note) reinitialize(){
	n.symbol.SetDict(axiom.TwelveNotes(n.tonalIndex())) 
	n.symbol.SetIndex(n.Index())
	n.degree.SetDict(axiom.Degrees(n.keyIndex()))
	n.degree.SetIndex(n.Index())
}
package tonality

import (
	"golang_music/internal/axiom"
	"golang_music/internal/util/sign"
)

// Tonality :
type Tonality struct {
	tonal sign.Sign
	key   sign.Sign
}

// New :
func New(tonal string, key string) *Tonality {
	t := new(Tonality)
	t.tonal = *sign.New(axiom.Tonal())
	t.SetTonalValue(tonal)
	t.key = *sign.New(axiom.TwelveNotes(t.TonalIndex()))
	t.SetKeyValue(key)
	return t
}

// TonalIndex :
func (t *Tonality) TonalIndex() uint {
	return t.tonal.Index()
}

// SetTonalIndex :
func (t *Tonality) SetTonalIndex(tonal uint) {
	t.tonal.SetIndex(tonal)
	t.reinitialize()
}

// TonalValue :
func (t *Tonality) TonalValue() string {
	return t.tonal.Value()
}

// SetTonalValue :
func (t *Tonality) SetTonalValue(tonal string) {
	t.tonal.SetValue(tonal)
	t.reinitialize()
}

// KeyIndex :
func (t *Tonality) KeyIndex() uint {
	return t.key.Index()
}

// SetKeyIndex :
func (t *Tonality) SetKeyIndex(key uint) {
	t.key.SetIndex(key)
}

// KeyValue :
func (t *Tonality) KeyValue() string {
	return t.key.Value()
}

// SetKeyValue :
func (t *Tonality) SetKeyValue(key string) {
	t.key.SetValue(key)
}

func (t *Tonality) reinitialize() {
	keyIndex := t.KeyIndex()
	t.key = *sign.New(axiom.TwelveNotes(t.TonalIndex()))
	t.SetKeyIndex(keyIndex)
}

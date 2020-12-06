package sign

import "golang_music/internal/util/dict"

// Sign :
type Sign struct {
	index uint
	value string
	seed  dict.Dict
}

// New :
func New(seed dict.Dict) *Sign {
	i := new(Sign)
	i.seed = seed
	return i
}

// Index : attribute
func (sign *Sign) Index() uint {
	return sign.index
}

// SetIndex : attribute
func (sign *Sign) SetIndex(index uint) {
	sign.index = index
	sign.value = sign.valueOf(index)
}

// Value : attribute
func (sign *Sign) Value() string {
	return sign.value
}

// SetValue : attribute
func (sign *Sign) SetValue(value string) {
	sign.value = value
	sign.index = sign.indexOf(value)
}

func (sign *Sign) valueOf(index uint) string {
	return sign.seed.ValueOf(index)
}

func (sign *Sign) indexOf(value string) uint {
	return sign.seed.IndexOf(value)
}

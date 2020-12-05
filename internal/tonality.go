package tonality

type Tonality struct {
	tonal tonal
	key   key
}

type tonal struct {
	index uint
	sign  string
}

type key struct {
	index uint
	sign  string
}

func (tonal *tonal) Sign() string {
	return tonal.sign
}

func (tonal *tonal) SetSign() string {
	return tonal.sign
}

func (tonal *tonal) Index() uint {
	return tonal.index
}

func (tonal *tonal) SetIndex() uint {
	return tonal.index
}

func (key *key) Sign() string {
	return key.sign
}

func (key *key) SetSign() string {
	return key.sign
}

func (key *key) Index() uint {
	return key.index
}

func (key *key) SetIndex() uint {
	return key.index
}

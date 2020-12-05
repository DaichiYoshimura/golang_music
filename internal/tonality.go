package tonality

// Tonality :
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

func (tonal *tonal) SetSign(sign string){
	tonal.sign = sign
}

func (tonal *tonal) Index() uint {
	return tonal.index
}

func (tonal *tonal) SetIndex(index uint){
	tonal.index = index
}

func (key *key) Sign() string {
	return key.sign
}

func (key *key) SetSign(sign string){
	key.sign = sign
}

func (key *key) Index() uint {
	return key.index
}

func (key *key) SetIndex(index uint){
	key.index = index
}

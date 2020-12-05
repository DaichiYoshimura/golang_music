package tonality

type Tonality struct {
	tonal tonal
	key   key
}

type tonal indexSign

type key indexSign
type indexSign struct {
	index uint
	sign string
}

func (indexSign *indexSign) Sign() string {
	return indexSign.sign
}

func (indexSign *indexSign) SetSign(sign string){
	indexSign.sign = sign
}

func (indexSign *indexSign) Index() uint {
	return indexSign.index
}

func (indexSign *indexSign) SetIndex(index uint){
	indexSign.index = index
}
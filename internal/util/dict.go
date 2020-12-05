package util

type Dict struct {
	index uint
	sign string
}

func (dict *Dict) Sign() string {
	return dict.sign
}

func (dict *Dict) SetSign(sign string){
	dict.sign = sign
}

func (dict *Dict) Index() uint {
	return dict.index
}

func (dict *Dict) SetIndex(index uint){
	dict.index = index
}
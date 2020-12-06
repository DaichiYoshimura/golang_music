package main

import "golang_music/internal/tonality"

func main() {

	t := tonality.New("major", "C")
	t.SetKeyIndex(2)
	print(t.KeyValue())
	t.SetTonalIndex(2)
	print(t.KeyValue())
}

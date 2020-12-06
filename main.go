package main

import "golang_music/internal/tonality"

func main() {

	t := tonality.New("major", "C")
	print(t.TonalValue())
}

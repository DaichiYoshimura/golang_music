package main

import (
	"golang_music/internal/note"
	"golang_music/internal/tonality"
	"strconv"
)

func main() {
	t := tonality.New("major", "C")
	n := note.New(t)
	n.SetSymbol("B")
	printAttr(*n)
	n.SetDegree("5.0")
	printAttr(*n)
	n.SetIndex(2)
	printAttr(*n)
	t.SetKeyValue("B-")
	printAttr(*n)
	t.SetKeyIndex(2)
	printAttr(*n)
	t.SetTonalValue("minor")
	printAttr(*n)
}

func printAttr(n note.Note){
	print(strconv.FormatUint(uint64(n.Index()), 10)+", "+string(n.Symbol())+", "+string(n.Degree())+"\n")
}

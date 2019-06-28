package main

import (
	"fmt"
	"sort"
)

var sharpScale = sort.StringSlice{"A", "A#", "B","C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var sharpStarter = sort.StringSlice{"G", "D", "A", "E", "B", "F#","e", "b", "f#", "c#", "g#", "d#"}
var flatScale = []string{"Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E","F", "Gb", "G"}

func main(){
	result := sharpStarter.Search("G")
	fmt.Println(result)

}
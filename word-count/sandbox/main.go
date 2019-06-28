package main

import (
	"fmt"
	"regexp"
)

type Frequency map[string]int


func WordCount(s string){
	a := regexp.MustCompile("[a-zA-Z0-9]+")
	fmt.Println(a.FindAllStringSubmatch(s,-1))
}

func main(){
	WordCount("ciao,-nana")
}
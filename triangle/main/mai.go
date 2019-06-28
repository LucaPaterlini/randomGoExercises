package main

import (
	"fmt"
	"math"
)

func main(){
	a := math.NaN()
	b := 1.0
	fmt.Println(math.IsNaN(a+b))
}

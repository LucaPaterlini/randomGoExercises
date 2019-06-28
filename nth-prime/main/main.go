package main

import (
	"fmt"
	"math"
)

func isPrime(x int)bool{
	for i:=2; i<=int(math.Sqrt(float64(x)));i++{
		if x%i==0 {return false}
	}
	return true
}

func Nth(n int)(int,bool){
	if n<1 {return 0, false}
	i:=1
	for n>0{
		i++
		if isPrime(i){n--}
	}
	return i,true
}

func main(){
	for i:=1;i<100;i++{
		if isPrime(i){
			fmt.Println(i)
		}
	}
	fmt.Println(Nth(1))
}
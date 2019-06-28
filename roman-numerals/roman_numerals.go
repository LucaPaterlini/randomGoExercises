// Package romannumerals implements a function to convert decimal to roman numbers
package romannumerals

import (
	"errors"
)


type pairDecNumeral struct {
	val int
	numeral string
}
var arrayDecNumeral = []pairDecNumeral{
	{1000,"M"},
	{900,"CM"},
	{500,"D"},
	{400,"CD"},
	{100,"C"},
	{90,  "XC"},
	{50, "L"},
	{40,"XL"},
	{10,"X"},
	{9,"IX"},
	{5,"V"},
	{4,"IV"},
	{1,"I"},
}


// ToRomanNumeral converts decimal to roman numeral
func ToRomanNumeral(x int) (romanS string, err error) {
	// check for negative numbers
	if x < 1  || x>3000{
		err = errors.New("out of bound! supported numbers 1 -3000")
	}else {
		for i:=0;i<13;i++{
			for x>=arrayDecNumeral[i].val{
				x -=arrayDecNumeral[i].val
				romanS+=arrayDecNumeral[i].numeral
			}
		}
	}
	return
}
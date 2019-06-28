// Package protein implements a set of functions that support the detection of proteins in rna
package protein

import (
	"errors"
)

var codonToProtein = map[string]string{
	"AUG":"Methionine",
	"UUU":"Phenylalanine",
	"UUC":"Phenylalanine",
	"UUA":"Leucine",
	"UUG":"Leucine",
	"UCU":"Serine",
	"UCC":"Serine",
	"UCA":"Serine",
	"UCG":"Serine",
	"UAU":"Tyrosine",
	"UAC":"Tyrosine",
	"UGU":"Cysteine",
	"UGC":"Cysteine",
	"UGG":"Tryptophan",
	"UAA":"",
	"UAG":"",
	"UGA":"",
}

var ErrStop  = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

// FromCodon converts a codon in its associated protein
func FromCodon(c string)(p string,err error){
	var ok bool
	if p,ok =codonToProtein[c];ok{
		if p==""{
			err = ErrStop
			return
		}
	}else{
		err = ErrInvalidBase
	}
	return
}

// FromRNA converts a rna in its array of proteins
func FromRNA(rna string)(p []string, e error){
	for i:=0;i<len(rna)-2;i+=3{
		if tmp,err := FromCodon(rna[i:i+3]);err==nil{
			p = append(p, tmp)
		}else {
			if err==ErrInvalidBase{
				e=err
			}
			return
		}
	}
	return
}
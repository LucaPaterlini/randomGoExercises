// Package bob contains the suite that calculate bob answers
package bob

import (
	"strings"
	"unicode"
)

func haveLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// check if something has been said
	remark = strings.TrimSpace(remark)
	if remark==""{
		return "Fine. Be that way!"
	}

	index := len(remark)-1
	question,yell := 0,0
	// check if the remark is a question
	if '?'==remark[index] {
		question=1
	}else{
		index-=1
	}
	// check if the remark is a yell
	if strings.ToUpper(remark)==remark && haveLetter(remark){
		yell=1
	}
	responseArray := []string{"Whatever.",
		"Whoa, chill out!","Sure.","Calm down, I know what I'm doing!"}

	return responseArray[2*question+yell]
}

// Package strain implements keep and discards for list of ints, list of strings and list of list of ints
package strain

type Ints []int
type Strings []string
type Lists [][]int

// Keep return a list of int that respect the condition checked using pred
func (list Ints) Keep(pred func(int) bool) (res Ints) {
	if list==nil{return nil}
	res = make(Ints, 0, len(list))
	for _, x := range list {
		if pred(x) {
			res = append(res, x)
		}
	}
	return
}

// Keep return a list of int that not respect the condition checked using pred
func (list Ints) Discard(pred func(int) bool) (res Ints) {
	if list==nil{return nil}
	res = make(Ints, 0, len(list))
	for _, x := range list {
		if !pred(x) {
			res = append(res, x)
		}
	}
	return
}

// Keep return a list of string that respect the condition checked using pred
func (list Strings) Keep(pred func(string) bool) (res Strings) {
	if list==nil{return nil}
	res = make(Strings, 0, len(list))
	for _, x := range list {
		if pred(x) {
			res = append(res, x)
		}
	}
	return
}

// Keep return a list of list of int that respect the condition checked using pred
func (list Lists) Keep(pred func([]int) bool) (res Lists) {
	if list==nil{return nil}
	res = make(Lists, 0, len(list))
	for _, x := range list {
		if pred(x) {
			res = append(res, x)
		}
	}
	return
}

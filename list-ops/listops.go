// Package listops provide a suite of struct and them methods
// to apply functions over arrays
package listops

type binFunc func(int, int) int
type IntList []int
type predFunc func(int) bool
type unaryFunc func(x int)int

// Foldr reduce the attribute array visiting from right to left and using
// the array item as first and the prev result as second parameter
func (L IntList) Foldr(f binFunc, i int) int {
	for j := len(L) - 1; j >= 0; j-- {
		i = f(L[j], i)
	}
	return i
}

// Foldl reduce the attribute array visiting from right to left and using
// the prev result as first parameter and the array item as second
func (L IntList) Foldl(f binFunc, i int) int {
	for j := len(L) - 1; j >= 0; j-- {
		i = f(i, L[j])
	}
	return i
}

// Filter apply a filter function on each element of the array
func (L IntList) Filter(f predFunc) (ret IntList) {
	ret = make([]int,L.Length())
	i:=0
	for _, x := range L {
		if f(x) {
			ret[i]= x
			i++
		}
	}
	ret = ret[:i]
	return
}

// Length calculate the len of the array
func (L IntList) Length() (c int) {
	for range L {
		c += 1
	}
	return
}
// Map apply a function to each element of the attribute array
func(L IntList)Map(f unaryFunc)(R IntList){
	R = make([]int,L.Length())
	var i int
	for i=0;i<L.Length();i++{
		R[i]=f(L[i])
	}
	return R[:i]
}
// Reverse returns the array with the positions of the items in reverse
func(L IntList)Reverse()(R IntList){
	R =make([]int,L.Length())
	for i:=0;i<L.Length();i++{
		R[L.Length()-1-i]=L[i]
	}
	return
}
// Append  retruns the array as parameter appended as the attribute array
func(L IntList)Append(x IntList)(res IntList){
	res = make([]int,x.Length()+L.Length())
	j:=0
	for _,y := range L{
			res[j]=y
			j++
	}
	for _,y := range x{
		res[j]=y
		j++
	}
	return
}
// Concat append all the int array inside the array if intList as parameter
// and returns a single IntList
func (L IntList)Concat(ll []IntList)(res IntList){
	res = L
	for _,x :=range ll{
		res =res.Append(x)
	}
	return
}

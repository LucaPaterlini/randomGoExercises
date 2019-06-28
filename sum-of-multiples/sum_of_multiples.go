// Package summultiples implements SumMultiples function.
package summultiples

// SumMultiples calculate the sum of all the unique multiples of particular
// numbers up to but not including that number.
func SumMultiples(limit int, divisors ...int)(res int){
	sieveArray := make([]bool,limit+1)
	for _,x := range divisors {
		if x==0 {continue}
		for i:=1;i*x<limit;i++{
			sieveArray[i*x]=true
		}
	}
	for i,b:= range sieveArray{
		if b{res+=i}
	}
	return
}

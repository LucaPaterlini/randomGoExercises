// Package allergies contains a suite of function about allergy score.
package allergies

type allergyRecord struct {
	score int
	name  string
}

var allergiesList = []allergyRecord{
	{0, "eggs"},
	{1, "peanuts"},
	{2, "shellfish"},
	{3, "strawberries"},
	{4, "tomatoes"},
	{5, "chocolate"},
	{6, "pollen"},
	{7, "cats"},
}

// Allergies returns the list of allergies for a given allergy score.
func Allergies(x uint) (res []string) {
	for i := 0; x != 0 && i < 8; i++ {
		if x%2 == 1 {
			res = append(res, allergiesList[i].name)
		}
		x /= 2
	}
	return
}

func in(arrayS []string, s string) bool {
	for _, item := range arrayS {
		if item == s {
			return true
		}
	}
	return false
}

// AllergicTo checks if an allergy is contained in those of the input score.
func AllergicTo(x uint, allergy string) bool {
	return in(Allergies(x), allergy)
}

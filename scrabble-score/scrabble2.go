package scrabble

var score = []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

func init() {
	score = append(score, append([]int{0, 0, 0, 0, 0, 0}, score...)...)
}

func Score2(word string) (result int) {
	for _, char := range word {
		result += score[char-'A']
	}
	return result
}


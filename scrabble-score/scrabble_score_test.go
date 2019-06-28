package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, test := range scrabbleScoreTests {
		if actual := Score(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}

func TestScore2(t *testing.T) {
	for _, test := range scrabbleScoreTests {
		if actual := Score2(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkScore2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score2(test.input)
		}
	}
}


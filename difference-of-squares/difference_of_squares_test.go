package diffsquares

import "testing"

var tests = []struct{ n, sqOfSum, sumOfSq float64 }{
	{5, 225, 55},
	{10, 3025, 385},
	{100, 25502500, 338350},
	{1000, 250500250000, 333833500},
}

func TestDifference(t *testing.T) {
	for _, test := range tests {
		want := test.sqOfSum - test.sumOfSq
		if s := Difference(test.n); s != want {
			t.Fatalf("Difference(%f) = %f, want %f", test.n, s, want)
		}
	}
}

func BenchmarkDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Difference(float64(i))
	}
}

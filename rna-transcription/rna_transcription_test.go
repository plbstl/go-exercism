package strand

import "testing"

func TestToRNA(t *testing.T) {
	for _, test := range rnaTests {
		actual, err := ToRNA(test.input)
		if test.expectedError && (err == nil) {
			t.Fatalf("FAIL: %s\nToRNA(%q)\nExpected error\nActual: %#v",
				test.description, test.input, actual)
		}
		if actual != test.expected {
			t.Fatalf("FAIL: %s - ToRNA(%q): %q, expected %q",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func BenchmarkToRNA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range rnaTests {
			ToRNA(test.input)
		}
	}
}

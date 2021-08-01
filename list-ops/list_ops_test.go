package listops

import (
	"reflect"
	"testing"
)

func TestIntList_Append(t *testing.T) {
	for _, tt := range appendTestCases {
		got := tt.list.Append(tt.appendThis)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

func TestIntList_Concat(t *testing.T) {
	for _, tt := range concatTestCases {
		got := tt.list.Concat(tt.args...)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

func TestIntList_Filter(t *testing.T) {
	for _, tt := range filterTestCases {
		in := IntList(tt.list)
		got := in.Filter(tt.fn)
		if !reflect.DeepEqual(IntList(tt.want), got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

func TestIntList_Length(t *testing.T) {
	for _, tt := range lengthTestCases {
		got := tt.list.Length()
		if tt.want != got {
			t.Fatalf("FAIL: %s: %q -- expected: %d, actual: %d", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

func TestIntList_Map(t *testing.T) {
	for _, tt := range mapTestCases {
		got := tt.list.Map(tt.fn)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

func TestIntList_Fold(t *testing.T) {
	var got int
	for _, tt := range foldTestCases {
		if tt.property == "foldr" {
			got = tt.list.Foldr(tt.fn, tt.initial)
		} else {
			got = tt.list.Foldl(tt.fn, tt.initial)
		}
		if got != tt.want {
			t.Fatalf("FAIL: %s: %q -- expected: %d, actual: %d", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}
func TestIntList_Reverse(t *testing.T) {
	for _, tt := range reverseTestCases {
		got := tt.list.Reverse()
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("FAIL: %s: %q -- expected: %v, actual: %v", tt.property, tt.name, tt.want, got)
		} else {
			t.Logf("PASS: %s: %s", tt.property, tt.name)
		}
	}
}

func BenchmarkIntList_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range appendTestCases {
			tt.list.Append(tt.appendThis)
		}
	}
}

func BenchmarkIntList_Concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range concatTestCases {
			tt.list.Concat(tt.args...)
		}
	}
}

func BenchmarkIntList_Filter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range filterTestCases {
			IntList(tt.list).Filter(tt.fn)
		}
	}
}

func BenchmarkIntList_Length(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range lengthTestCases {
			tt.list.Length()
		}
	}
}

func BenchmarkIntList_Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range mapTestCases {
			tt.list.Map(tt.fn)
		}
	}
}

func BenchmarkIntList_Fold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range foldTestCases {
			if tt.property == "foldr" {
				tt.list.Foldr(tt.fn, tt.initial)
			} else {
				tt.list.Foldl(tt.fn, tt.initial)
			}
		}
	}
}
func BenchmarkIntList_Reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range reverseTestCases {
			tt.list.Reverse()
		}
	}
}

package strain

import (
	"reflect"
	"testing"
)

func TestInts_Keep(t *testing.T) {
	for _, test := range intsTests {
		// setup here copies test.list, preserving the nil value if it is nil
		// and making a fresh copy of the underlying array otherwise.
		cp := test.list
		if cp != nil {
			cp = append(Ints{}, cp...)
		}
		switch res := cp.Keep(test.pred); {
		case !reflect.DeepEqual(cp, test.list):
			t.Fatalf("%#v.Keep() should not modify its receiver.  "+
				"Found %#v, receiver should stay %#v",
				test.list, cp, test.list)
		case !reflect.DeepEqual(res, test.wantKeep):
			t.Fatalf("%#v.Keep()\ngot: %#v\nwant: %#v",
				test.list, res, test.wantKeep)
		}
	}
}

func TestInts_Discard(t *testing.T) {
	for _, test := range intsTests {
		cp := test.list
		if cp != nil {
			cp = append(Ints{}, cp...) // dup underlying array
		}
		switch res := cp.Discard(test.pred); {
		case !reflect.DeepEqual(cp, test.list):
			t.Fatalf("%#v.Discard() should not modify its receiver.  "+
				"Found %#v, receiver should stay %#v",
				test.list, cp, test.list)
		case !reflect.DeepEqual(res, test.wantDiscard):
			t.Fatalf("%#v.Discard()\ngot: %#v\nwant: %#v",
				test.list, res, test.wantDiscard)
		}
	}
}

func TestLists_Keep(t *testing.T) {
	cp := append(Lists{}, listsTest.list...)
	switch res := cp.Keep(listsTest.pred); {
	case !reflect.DeepEqual(cp, listsTest.list):
		t.Fatalf("%#v.Keep() should not modify its receiver.  "+
			"Found %#v, receiver should stay %#v",
			listsTest.list, cp, listsTest.list)
	case !reflect.DeepEqual(res, listsTest.wantKeep):
		t.Fatalf("%#v.Keep()\ngot: %#v\nwant: %#v", listsTest.list, res, listsTest.wantKeep)
	}
}

func TestLists_Discard(t *testing.T) {
	cp := append(Lists{}, listsTest.list...)
	switch res := cp.Discard(has5); {
	case !reflect.DeepEqual(cp, listsTest.list):
		t.Fatalf("%#v.Discard() should not modify its receiver.  "+
			"Found %#v, receiver should stay %#v",
			listsTest.list, cp, listsTest.list)
	case !reflect.DeepEqual(res, listsTest.wantDiscard):
		t.Fatalf("%#v.Discard()\ngot: %#v\nwant: %#v", listsTest.list, res, listsTest.wantDiscard)
	}
}

func TestStrings_Keep(t *testing.T) {
	cp := append(Strings{}, stringsTest.list...) // make copy, as with TestInts
	switch res := cp.Keep(stringsTest.pred); {
	case !reflect.DeepEqual(cp, stringsTest.list):
		t.Fatalf("%#v.Keep() should not modify its receiver.  "+
			"Found %#v, receiver should stay %#v",
			stringsTest.list, cp, stringsTest.list)
	case !reflect.DeepEqual(res, stringsTest.wantKeep):
		t.Fatalf("%#v.Keep()\ngot: %#v\nwant: %#v", stringsTest.list, res, stringsTest.wantKeep)
	}
}

func TestStrings_Discard(t *testing.T) {
	cp := append(Strings{}, stringsTest.list...) // make copy, as with TestInts
	switch res := cp.Discard(stringsTest.pred); {
	case !reflect.DeepEqual(cp, stringsTest.list):
		t.Fatalf("%#v.Discard() should not modify its receiver.  "+
			"Found %#v, receiver should stay %#v",
			stringsTest.list, cp, stringsTest.list)
	case !reflect.DeepEqual(res, stringsTest.wantDiscard):
		t.Fatalf("%#v.Discard()\ngot: %#v\nwant: %#v", stringsTest.list, res, stringsTest.wantDiscard)
	}
}

func BenchmarkInts_Keep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range intsTests {
			test.list.Keep(test.pred)
		}
	}
}

func BenchmarkInts_Discard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range intsTests {
			test.list.Discard(test.pred)
		}
	}
}

func BenchmarkLists_Keep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := listsTest.list
		l.Keep(listsTest.pred)
	}
}

func BenchmarkLists_Discard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := listsTest.list
		l.Discard(listsTest.pred)
	}
}

func BenchmarkStrings_Keep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stringsTest.list
		s.Keep(stringsTest.pred)
	}
}

func BenchmarkStrings_Discard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stringsTest.list
		s.Discard(stringsTest.pred)
	}
}

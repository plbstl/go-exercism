package strain

func lt10(x int) bool { return x < 10 }
func gt10(x int) bool { return x > 10 }
func odd(x int) bool  { return x&1 == 1 }
func even(x int) bool { return x&1 == 0 }

var intsTests = []struct {
	pred        func(int) bool
	list        Ints
	wantKeep    Ints
	wantDiscard Ints
}{
	{
		pred:        lt10,
		list:        nil,
		wantKeep:    nil,
		wantDiscard: nil},
	{
		pred:        lt10,
		list:        Ints{1, 2, 3},
		wantKeep:    Ints{1, 2, 3},
		wantDiscard: nil},
	{
		pred:        gt10,
		list:        Ints{1, 2, 3},
		wantKeep:    nil,
		wantDiscard: Ints{1, 2, 3}},
	{
		pred:        odd,
		list:        Ints{1, 2, 3},
		wantKeep:    Ints{1, 3},
		wantDiscard: Ints{2}},
	{
		pred:        even,
		list:        Ints{1, 2, 3, 4, 5},
		wantKeep:    Ints{2, 4},
		wantDiscard: Ints{1, 3, 5}},
}

func has5(l []int) bool {
	for _, e := range l {
		if e == 5 {
			return true
		}
	}
	return false
}

var listsTest = struct {
	pred        func([]int) bool
	list        Lists
	wantKeep    Lists
	wantDiscard Lists
}{
	pred: has5,
	list: Lists{
		{1, 2, 3},
		{5, 5, 5},
		{5, 1, 2},
		{2, 1, 2},
		{1, 5, 2},
		{2, 2, 1},
		{1, 2, 5},
	},
	wantKeep: Lists{
		{5, 5, 5},
		{5, 1, 2},
		{1, 5, 2},
		{1, 2, 5},
	},
	wantDiscard: Lists{
		{1, 2, 3},
		{2, 1, 2},
		{2, 2, 1},
	},
}

func zword(s string) bool { return len(s) > 0 && s[0] == 'z' }

var stringsTest = struct {
	pred        func(string) bool
	list        Strings
	wantKeep    Strings
	wantDiscard Strings
}{
	pred:        zword,
	list:        Strings{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"},
	wantKeep:    Strings{"zebra", "zombies", "zealot"},
	wantDiscard: Strings{"apple", "banana", "cherimoya"},
}

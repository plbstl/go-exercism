package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

var appendTestCases = []struct {
	name       string
	property   string
	list       IntList
	appendThis IntList
	want       IntList
}{
	{
		name:       "empty list",
		property:   "append",
		list:       []int{},
		appendThis: []int{},
		want:       []int{},
	},
	{
		name:       "empty list to list",
		property:   "append",
		list:       []int{},
		appendThis: []int{1, 2, 3, 4},
		want:       []int{1, 2, 3, 4},
	},
	{
		name:       "non-empty lists",
		property:   "append",
		list:       []int{1, 2},
		appendThis: []int{2, 3, 4, 5},
		want:       []int{1, 2, 2, 3, 4, 5},
	},
}

var concatTestCases = []struct {
	name     string
	property string
	list     IntList
	args     []IntList
	want     IntList
}{
	{
		name:     "empty list",
		property: "concat",
		list:     []int{},
		args:     []IntList{},
		want:     []int{},
	},
	{
		name:     "list of lists",
		property: "concat",
		list:     []int{1, 2},
		args:     []IntList{[]int{3}, []int{}, []int{4, 5, 6}},
		want:     []int{1, 2, 3, 4, 5, 6},
	},
}

var filterTestCases = []struct {
	name     string
	property string
	fn       predFunc
	list     []int
	want     []int
}{
	{
		name:     "empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{1, 2, 3, 4, 5},
		want:     []int{1, 3, 5},
	},
}

var lengthTestCases = []struct {
	name     string
	property string
	list     IntList
	want     int
}{
	{
		name:     "empty list",
		property: "length",
		list:     []int{},
		want:     0,
	},
	{
		name:     "non-empty list",
		property: "length",
		list:     []int{1, 2, 3, 4},
		want:     4,
	},
}

var mapTestCases = []struct {
	name     string
	property string
	list     IntList
	fn       unaryFunc
	want     IntList
}{
	{
		name:     "empty list",
		property: "map",
		list:     []int{},
		fn:       func(x int) int { return x + 1 },
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "map",
		list:     []int{1, 3, 5, 7},
		fn:       func(x int) int { return x + 1 },
		want:     []int{2, 4, 6, 8},
	},
}

var foldTestCases = []struct {
	name     string
	property string
	fn       binFunc
	initial  int
	list     IntList
	want     int
}{
	{
		name:     "empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     0,
		list:     []int{2, 5},
	},
	{
		name:     "empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     2,
		list:     []int{2, 5},
	},
}

var reverseTestCases = []struct {
	name     string
	property string
	list     IntList
	want     IntList
}{
	{
		name:     "empty list",
		property: "reverse",
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "reverse",
		list:     []int{1, 3, 5, 7},
		want:     []int{7, 5, 3, 1},
	},
}

package strand

// Source: exercism/problem-specifications
// Commit: 294c831 rna-transcription: add test for empty sequence (#1266)
// Problem Specifications Version: 1.3.0

var rnaTests = []struct {
	description   string
	input         string
	expected      string
	expectedError bool
}{
	{"Empty RNA sequence",
		"", "",
		false,
	},
	{"RNA complement of cytosine is guanine",
		"C", "G",
		false,
	},
	{"RNA complement of guanine is cytosine",
		"G", "C",
		false,
	},
	{"RNA complement of thymine is adenine",
		"T", "A",
		false,
	},
	{"RNA complement of adenine is uracil",
		"A", "U",
		false,
	},
	{"RNA complement",
		"ACGTGGTCTTAA", "UGCACCAGAAUU",
		false,
	},
	{"RNA complement with error of lowercase letter",
		"oTTAA", "",
		true,
	},
	{"RNA complement with error of invalid letter",
		"U", "",
		true,
	},
}

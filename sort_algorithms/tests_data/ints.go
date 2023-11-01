package tests_data

var SortTestsInt = []struct {
	Name string
	Args args[int]
	Want []int
}{
	{
		Name: "Random elements",
		Args: args[int]{Input: []int{
			11, 4, 3, 6, 7, 2, 5, 1, 9, 8,
		}},
		Want: []int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 11,
		},
	},
	{
		Name: "Single element",
		Args: args[int]{Input: []int{11}},
		Want: []int{11},
	},
	{
		Name: "Random elements with negatives",
		Args: args[int]{Input: []int{
			4, 3, 6, -7, 2, 5, 1, 9, 8, -11,
		}},
		Want: []int{
			-11, -7, 1, 2, 3, 4, 5, 6, 8, 9,
		},
	},
	{
		Name: "Random elements with duplicates",
		Args: args[int]{Input: []int{
			7, 11, 4, 3, 6, 7, 2, 5, 1, 9, 8, 2,
		}},
		Want: []int{
			1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9, 11,
		},
	},
	{
		Name: "Random elements with negatives and duplicates",
		Args: args[int]{Input: []int{
			2, 4, 3, -3, -7, 6, -7, 2, 5, 1, 9, 8, -11,
		}},
		Want: []int{
			-11, -7, -7, -3, 1, 2, 2, 3, 4, 5, 6, 8, 9,
		},
	},
}

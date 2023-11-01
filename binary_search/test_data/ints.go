package test_data

var SortTestsInt = []struct {
	Name string
	Args args[int]
	Want int
}{
	{
		Name: "Random elements",
		Args: args[int]{Input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11},
			SearchValue: 5,
		},
		Want: 4,
	},
	{
		Name: "Single element (not found)",
		Args: args[int]{Input: []int{11}, SearchValue: 0},
		Want: -1,
	},
	{
		Name: "Single element (found)",
		Args: args[int]{Input: []int{11}, SearchValue: 11},
		Want: 0,
	},
	{
		Name: "Random elements with negatives",
		Args: args[int]{Input: []int{-11, -7, 1, 2, 3, 4, 5, 6, 8, 9}, SearchValue: -7},
		Want: 1,
	},
	{
		Name: "Random elements with duplicates",
		Args: args[int]{Input: []int{1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9, 11}, SearchValue: 8},
		Want: 9,
	},
	{
		Name: "Random elements with negatives and duplicates",
		Args: args[int]{Input: []int{-11, -7, -7, -3, 1, 2, 2, 3, 4, 5, 6, 8, 9}, SearchValue: 2},
		Want: 6,
	},
}

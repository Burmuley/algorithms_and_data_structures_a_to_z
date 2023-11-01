package tests_data

var SortTestsFloat = []struct {
	Name string
	Args args[float64]
	Want []float64
}{
	{
		Name: "Random elements",
		Args: args[float64]{Input: []float64{
			11.0, 127.365, 3.1, 16.0, 7, 2, 5, 1.2, 90, 8.1,
		}},
		Want: []float64{
			1.2, 2, 3.1, 5, 7, 8.1, 11.0, 16.0, 90, 127.365,
		},
	},
	{
		Name: "Single element",
		Args: args[float64]{Input: []float64{11}},
		Want: []float64{11},
	},
	{
		Name: "Random elements with negatives",
		Args: args[float64]{Input: []float64{
			4, 3, 6, -7, 2, 5, 1, 9, 8, -11,
		}},
		Want: []float64{
			-11, -7, 1, 2, 3, 4, 5, 6, 8, 9,
		},
	},
	{
		Name: "Random elements with duplicates",
		Args: args[float64]{Input: []float64{
			7, 11, 4, 3, 6, 7, 2, 5, 1, 9, 8, 2,
		}},
		Want: []float64{
			1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 9, 11,
		},
	},
	{
		Name: "Random elements with negatives and duplicates",
		Args: args[float64]{Input: []float64{
			2, 4, 3, -3, -7, 6, -7, 2, 5, 1, 9, 8, -11,
		}},
		Want: []float64{
			-11, -7, -7, -3, 1, 2, 2, 3, 4, 5, 6, 8, 9,
		},
	},
}

package merge

import (
	"github.com/Burmuley/algorithms_and_data_structures_a_to_z/sort_algorithms/tests_data"
	"reflect"
	"testing"
)

func TestSortInt(t *testing.T) {
	tests := tests_data.SortTestsInt
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			if got := Sort(tt.Args.Input); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Sort() = %v, want %v", got, tt.Want)
			} else {
				t.Log(got)
			}
		})
	}
}

func TestSortFloat(t *testing.T) {
	tests := tests_data.SortTestsFloat
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			if got := Sort(tt.Args.Input); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Sort() = %v, want %v", got, tt.Want)
			} else {
				t.Log(got)
			}
		})
	}
}

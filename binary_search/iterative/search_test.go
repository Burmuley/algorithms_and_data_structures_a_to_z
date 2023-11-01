package iterative

import (
	"github.com/Burmuley/algorithms_and_data_structures_a_to_z/binary_search/test_data"
	"reflect"
	"testing"
)

func TestBinarySearch_Int(t *testing.T) {
	tests := test_data.SortTestsInt
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			if got := BinarySearch(tt.Args.Input, tt.Args.SearchValue); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.Want)
			} else {
				t.Log(got)
			}
		})
	}
}

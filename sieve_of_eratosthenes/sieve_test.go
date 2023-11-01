package sieve_of_eratosthenes

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "primes #1",
			args: args{max: 10},
			want: []int{2, 3, 5, 7},
		},
		{
			name: "primes #2",
			args: args{max: 30},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sieve(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

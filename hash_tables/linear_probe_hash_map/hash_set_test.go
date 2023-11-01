package linear_probe_hash_map

import (
	"cmp"
	"reflect"
	"testing"
)

func Test_murmurHash(t *testing.T) {
	type args[K comparable] struct {
		v K
		c uint64
	}
	type testCase[K comparable] struct {
		name  string
		args  args[K]
		iters int
	}
	tests := []testCase[int]{
		{
			name:  "Less than capacity",
			args:  args[int]{v: 5, c: 10},
			iters: 10,
		},
		{
			name:  "Same as capacity",
			args:  args[int]{v: 10, c: 10},
			iters: 20,
		},
		{
			name:  "Larger than capacity",
			args:  args[int]{v: 555, c: 10},
			iters: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := make([]uint64, 0, tt.iters)
			for i := 0; i < tt.iters; i++ {
				res := murmurHash(tt.args.v, tt.args.c)
				results = append(results, res)
				if res >= tt.args.c {
					t.Errorf("internalHash() provided hash larger than capacity: %v", results)
				}
			}
			for i := 1; i < len(results); i++ {
				if results[i] != results[i-1] {
					t.Errorf("internalHash() provided incosistent hashes: %v", results)
				}
			}
		})
	}
}

func TestHashSet_Count(t *testing.T) {
	type testCase[K cmp.Ordered, V any] struct {
		name string
		hs   *HashSet[K, V]
		want int
	}
	tests := []testCase[int, int]{
		{
			name: "empty hash",
			hs:   NewHashSet[int, int](),
			want: 0,
		},
		{
			name: "single element hash",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			want: 1,
		},
		{
			name: "multiple elements hash",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hs.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Get(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name     string
		hs       *HashSet[K, V]
		args     args[K]
		want     V
		wantCond bool
	}
	tests := []testCase[int, int]{
		{
			name:     "empty hash",
			hs:       NewHashSet[int, int](),
			args:     args[int]{key: 10},
			want:     0,
			wantCond: false,
		},
		{
			name: "single element hash (found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			args:     args[int]{key: 10},
			want:     50,
			wantCond: true,
		},
		{
			name: "single element hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			args:     args[int]{key: 100},
			want:     0,
			wantCond: false,
		},
		{
			name: "multiple elements hash (found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args:     args[int]{key: 10},
			want:     100,
			wantCond: true,
		},
		{
			name: "multiple elements hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args:     args[int]{key: 100},
			want:     0,
			wantCond: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.hs.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantCond {
				t.Errorf("Get() got condition = %v, want %v", got1, tt.wantCond)
			}
		})
	}
}

func TestHashSet_Add(t *testing.T) {
	type args[K cmp.Ordered, V any] struct {
		key   K
		value V
	}
	type testCase[K cmp.Ordered, V any] struct {
		name      string
		hs        *HashSet[K, V]
		args      args[K, V]
		wantCount int
	}
	tests := []testCase[int, int]{
		{
			name:      "empty hash",
			hs:        NewHashSet[int, int](),
			args:      args[int, int]{key: 10, value: 10},
			wantCount: 1,
		},
		{
			name: "single element hash",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			args:      args[int, int]{key: 20, value: 100},
			wantCount: 2,
		},
		{
			name: "multiple elements hash",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args:      args[int, int]{key: 200, value: 1000},
			wantCount: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hs.Add(tt.args.key, tt.args.value)
			if tt.hs.Count() != tt.wantCount {
				t.Errorf("Add() resulted in wrong number of elements = %v, want %v", tt.hs.Count(), tt.wantCount)
			}
			if !tt.hs.Contains(tt.args.key) {
				t.Errorf("Add() element not found after it was added, want Contains(%v)==true", tt.args)
			}
			if res, ok := tt.hs.Get(tt.args.key); res != tt.args.value || !ok {
				t.Errorf("Add() element added with wrong value, got = %v, want %v", res, tt.args.value)
			}
		})
	}
}

func TestHashSet_Contains(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name string
		hs   *HashSet[K, V]
		args args[K]
		want bool
	}
	tests := []testCase[int, int]{
		{
			name: "empty hash",
			hs:   NewHashSet[int, int](),
			args: args[int]{key: 10},
			want: false,
		},
		{
			name: "single element hash (found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			args: args[int]{key: 10},
			want: true,
		},
		{
			name: "single element hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			args: args[int]{key: 100},
			want: false,
		},
		{
			name: "multiple elements hash (found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args: args[int]{key: 10},
			want: true,
		},
		{
			name: "multiple elements hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args: args[int]{key: 100},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hs.Contains(tt.args.key); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Remove(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name      string
		hs        *HashSet[K, V]
		args      args[K]
		want      bool
		wantCount int
	}
	tests := []testCase[int, int]{
		{
			name:      "empty hash",
			hs:        NewHashSet[int, int](),
			args:      args[int]{key: 10},
			wantCount: 0,
			want:      false,
		},
		{
			name: "single element hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			args:      args[int]{key: 20},
			wantCount: 1,
			want:      false,
		},
		{
			name: "single element hash (found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 50)
				return hs
			}(),
			args:      args[int]{key: 10},
			wantCount: 0,
			want:      true,
		},
		{
			name: "multiple elements hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args:      args[int]{key: 200},
			wantCount: 20,
			want:      false,
		},
		{
			name: "multiple elements hash (found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args:      args[int]{key: 3},
			wantCount: 19,
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hs.Remove(tt.args.key); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Resize(t *testing.T) {
	type args struct {
		capacity int
	}
	type testCase[K cmp.Ordered, V any] struct {
		name      string
		hs        *HashSet[K, V]
		args      args
		wantKey   int
		wantCond  bool
		wantCount int
	}
	tests := []testCase[int, int]{
		{
			name:      "empty hash",
			hs:        NewHashSet[int, int](),
			args:      args{22},
			wantKey:   20,
			wantCond:  false,
			wantCount: 0,
		},
		{
			name: "single element hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 100)

				return hs
			}(),
			args:      args{22},
			wantKey:   20,
			wantCond:  false,
			wantCount: 1,
		},
		{
			name: "single element hash (found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				hs.Add(10, 100)

				return hs
			}(),
			args:      args{22},
			wantKey:   10,
			wantCond:  true,
			wantCount: 1,
		},
		{
			name: "multiple elements hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args:      args{44},
			wantKey:   20,
			wantCond:  false,
			wantCount: 20,
		},
		{
			name: "multiple elements hash (not found)",
			hs: func() *HashSet[int, int] {
				hs := NewHashSet[int, int]()
				for i := 0; i < 20; i++ {
					hs.Add(i, i*10)
				}

				return hs
			}(),
			args:      args{44},
			wantKey:   19,
			wantCond:  true,
			wantCount: 20,
		},
	}
	//prime := NewPrime()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hs.Resize(tt.args.capacity)
			//if res := tt.hs.capacity; res != int(prime.getPrime(uint(tt.args.capacity))) {
			if res := tt.hs.capacity; res != tt.args.capacity {
				t.Errorf("Resize() wrong capacity after resize = %v, want %v", res, tt.args.capacity)
			}
			if res := tt.hs.Count(); res != tt.wantCount {
				t.Errorf("Resize() wrong count after resize = %v, want %v", res, tt.wantCount)
			}
			if res := tt.hs.Contains(tt.wantKey); res != tt.wantCond {
				t.Errorf("Resize() bad key condition after resize = key %v exists %v, want key %v exists %v",
					tt.wantKey, tt.wantCond, tt.wantKey, res)
			}
		})
	}
}

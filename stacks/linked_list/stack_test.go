package linked_list

import (
	"errors"
	"reflect"
	"testing"
)

func TestStack_Count(t *testing.T) {
	type testCase[T any] struct {
		name  string
		stack *Stack[T]
		want  int
	}
	tests := []testCase[int]{
		{
			name:  "Empty stack",
			stack: NewStack[int](),
			want:  0,
		},
		{
			name: "Single element",
			stack: func() *Stack[int] {
				stack := NewStack[int]()
				stack.Push(10)
				return stack
			}(),
			want: 1,
		},
		{
			name: "Twenty elements",
			stack: func() *Stack[int] {
				stack := NewStack[int]()
				for i := 1; i <= 20; i++ {
					stack.Push(i)
				}

				return stack
			}(),
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type testCase[T any] struct {
		name    string
		stack   *Stack[T]
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name:    "Empty stack",
			stack:   NewStack[int](),
			want:    0,
			wantErr: ErrEmptyStack,
		},
		{
			name: "Single element",
			stack: func() *Stack[int] {
				stack := NewStack[int]()
				stack.Push(10)
				return stack
			}(),
			want:    10,
			wantErr: nil,
		},
		{
			name: "Twenty elements",
			stack: func() *Stack[int] {
				stack := NewStack[int]()
				for i := 1; i <= 20; i++ {
					stack.Push(i)
				}

				return stack
			}(),
			want:    20,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.stack.Peek(); !reflect.DeepEqual(got, tt.want) || !errors.Is(err, tt.wantErr) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type testCase[T any] struct {
		name    string
		stack   *Stack[T]
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name:    "Empty stack",
			stack:   NewStack[int](),
			want:    0,
			wantErr: ErrEmptyStack,
		},
		{
			name: "Single element",
			stack: func() *Stack[int] {
				stack := NewStack[int]()
				stack.Push(10)
				return stack
			}(),
			want:    10,
			wantErr: nil,
		},
		{
			name: "Twenty elements",
			stack: func() *Stack[int] {
				stack := NewStack[int]()
				for i := 1; i <= 20; i++ {
					stack.Push(i)
				}

				return stack
			}(),
			want:    20,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.stack.Pop(); !reflect.DeepEqual(got, tt.want) || !errors.Is(err, tt.wantErr) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name  string
		stack *Stack[T]
		want  *Stack[T]
		args  args[T]
	}
	tests := []testCase[int]{
		{
			name:  "To empty stack",
			stack: NewStack[int](),
			want: func() *Stack[int] {
				stack := NewStack[int]()
				stack.Push(10)
				return stack
			}(),
			args: struct{ value int }{value: 10},
		},
		{
			name: "To non-empty stack",
			stack: func() *Stack[int] {
				stack := NewStack[int]()
				for i := 0; i < 20; i++ {
					stack.Push(i)
				}
				return stack
			}(),
			want: func() *Stack[int] {
				stack := NewStack[int]()
				for i := 0; i < 20; i++ {
					stack.Push(i)
				}
				stack.Push(100)
				return stack
			}(),
			args: struct{ value int }{value: 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.stack.Push(tt.args.value); !reflect.DeepEqual(tt.stack, tt.want) {
				t.Errorf("Pop() = %v, want %v", tt.stack, tt.want)
			}
		})
	}
}

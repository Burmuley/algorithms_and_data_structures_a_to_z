package dummy

import (
	"errors"
	"reflect"
	"testing"
)

func TestQueue_Count(t *testing.T) {
	type testCase[T any] struct {
		name  string
		queue *Queue[T]
		want  int
	}
	tests := []testCase[int]{
		{
			name:  "Empty queue",
			queue: NewQueue[int](),
			want:  0,
		},
		{
			name: "Single element",
			queue: func() *Queue[int] {
				queue := NewQueue[int]()
				queue.Enqueue(10)
				return queue
			}(),
			want: 1,
		},
		{
			name: "Twenty elements",
			queue: func() *Queue[int] {
				queue := NewQueue[int]()
				for i := 1; i <= 20; i++ {
					queue.Enqueue(i)
				}

				return queue
			}(),
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.queue.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type testCase[T any] struct {
		name    string
		queue   *Queue[T]
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name:    "Empty queue",
			queue:   NewQueue[int](),
			want:    0,
			wantErr: ErrEmptyQueue,
		},
		{
			name: "Single element",
			queue: func() *Queue[int] {
				queue := NewQueue[int]()
				queue.Enqueue(10)
				return queue
			}(),
			want:    10,
			wantErr: nil,
		},
		{
			name: "Twenty elements",
			queue: func() *Queue[int] {
				queue := NewQueue[int]()
				for i := 1; i <= 20; i++ {
					queue.Enqueue(i)
				}

				return queue
			}(),
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.queue.Peek(); !reflect.DeepEqual(got, tt.want) || !errors.Is(err, tt.wantErr) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type testCase[T any] struct {
		name    string
		queue   *Queue[T]
		want    *Queue[T]
		wantErr error
	}
	tests := []testCase[int]{
		{
			name:    "Empty queue",
			queue:   NewQueue[int](),
			want:    NewQueue[int](),
			wantErr: ErrEmptyQueue,
		},
		{
			name: "Single element",
			queue: func() *Queue[int] {
				queue := NewQueue[int]()
				queue.Enqueue(10)
				return queue
			}(),
			want:    NewQueue[int](),
			wantErr: nil,
		},
		{
			name: "Twenty elements",
			queue: func() *Queue[int] {
				queue := NewQueue[int]()
				for i := 1; i <= 20; i++ {
					queue.Enqueue(i)
				}

				return queue
			}(),
			want: func() *Queue[int] {
				queue := NewQueueWithCapacity[int](20)
				for i := 1; i <= 20; i++ {
					queue.Enqueue(i)
				}
				_ = queue.Dequeue()
				return queue
			}(),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.queue.Dequeue(); !reflect.DeepEqual(tt.queue, tt.want) || !errors.Is(err, tt.wantErr) {
				t.Errorf("Pop() = %v, want %v", tt.queue, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name  string
		queue *Queue[T]
		want  *Queue[T]
		args  args[T]
	}
	tests := []testCase[int]{
		{
			name:  "To empty queue",
			queue: NewQueue[int](),
			want: func() *Queue[int] {
				queue := NewQueue[int]()
				queue.Enqueue(10)
				return queue
			}(),
			args: struct{ value int }{value: 10},
		},
		{
			name: "To non-empty queue",
			queue: func() *Queue[int] {
				queue := NewQueue[int]()
				for i := 0; i < 20; i++ {
					queue.Enqueue(i)
				}
				return queue
			}(),
			want: func() *Queue[int] {
				queue := NewQueue[int]()
				for i := 0; i < 20; i++ {
					queue.Enqueue(i)
				}
				queue.Enqueue(100)
				return queue
			}(),
			args: struct{ value int }{value: 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.queue.Enqueue(tt.args.value); !reflect.DeepEqual(tt.queue, tt.want) {
				t.Errorf("Pop() = %v, want %v", tt.queue, tt.want)
			}
		})
	}
}

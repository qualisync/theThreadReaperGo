package goroutines

import (
	"reflect"
	"sync"
	"testing"
)

func TestProcessNumberList(t *testing.T) {
	tests := []struct {
		name     string
		expected []int
	}{
		{
			name:     "Default numbers doubled",
			expected: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessNumberList(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s failed: expected %v, got %v", tt.name, tt.expected, got)
			}
		})
	}
}

func Test_doubleNumberToArray(t *testing.T) {
	type args struct {
		wg          *sync.WaitGroup
		lock        *sync.Mutex
		destination *int
		number      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Double 1",
			args: args{
				wg:          &sync.WaitGroup{},
				lock:        &sync.Mutex{},
				destination: new(int),
				number:      1,
			},
			want: 2,
		},
		{
			name: "Double 5",
			args: args{
				wg:          &sync.WaitGroup{},
				lock:        &sync.Mutex{},
				destination: new(int),
				number:      5,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		tt.args.wg.Add(1) // Setup WaitGroup
		t.Run(tt.name, func(t *testing.T) {
			doubleNumberToArray(tt.args.wg, tt.args.lock, tt.args.destination, tt.args.number)
			tt.args.wg.Wait() // Wait for the goroutine to complete

			if *tt.args.destination != tt.want {
				t.Errorf("%s failed: expected %v, got %v", tt.name, tt.want, *tt.args.destination)
			}
		})
	}
}

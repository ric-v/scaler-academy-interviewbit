package distinctnumbersinwindow

import (
	"reflect"
	"testing"
)

func TestDistinctNumbersInWindow(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 1, 3, 4, 3}, 3}, []int{2, 3, 3, 2}},
		{"test2", args{[]int{1, 1, 2, 2}, 1}, []int{1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctNumbersInWindow(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctNumbersInWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkDistinctNumbersInWindow-8   	 7160967	       170.3 ns/op	      32 B/op	       1 allocs/op
func BenchmarkDistinctNumbersInWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DistinctNumbersInWindow([]int{1, 2, 1, 3, 4, 3}, 3)
	}
}

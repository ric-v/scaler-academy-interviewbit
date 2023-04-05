package sumofoddindices

import (
	"reflect"
	"testing"
)

func TestSumOfOddIndices(t *testing.T) {
	type args struct {
		A []int
		B [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, [][]int{{0, 2}, {1, 4}}}, []int{2, 6}},
		{"test2", args{[]int{2, 1, 8, 3, 9}, [][]int{{0, 3}, {2, 4}}}, []int{4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfOddIndices(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfOddIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSumOfOddIndices-8   	19788978	        60.00 ns/op	      64 B/op	       2 allocs/op
func BenchmarkSumOfOddIndices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfOddIndices([]int{1, 2, 3, 4, 5}, [][]int{{0, 2}, {1, 4}})
	}
}

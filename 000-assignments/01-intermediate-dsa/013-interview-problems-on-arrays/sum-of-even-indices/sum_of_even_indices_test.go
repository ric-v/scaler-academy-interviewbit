package sumofevenindices

import (
	"reflect"
	"testing"
)

func TestSumOfEvenIndices(t *testing.T) {
	type args struct {
		A []int
		B [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, [][]int{{0, 3}, {2, 4}}}, []int{4, 8}},
		{"test2", args{[]int{2, 1, 8, 3, 9}, [][]int{{0, 3}, {2, 4}}}, []int{10, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfEvenIndices(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfEvenIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSumOfEvenIndices-8   	14922124	        68.78 ns/op	      64 B/op	       2 allocs/op
func BenchmarkSumOfEvenIndices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfEvenIndices([]int{1, 2, 3, 4, 5}, [][]int{{0, 3}, {2, 4}})
	}
}

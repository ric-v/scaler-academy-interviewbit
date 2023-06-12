package submatrixsumqueries

import (
	"reflect"
	"testing"
)

func TestSubMatrixSumQueries(t *testing.T) {
	type args struct {
		A [][]int
		B []int
		C []int
		D []int
		E []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2}, []int{1, 2}, []int{2, 3}, []int{2, 3}}, []int{12, 28}},
		{"test2", args{[][]int{{5, 17, 100, 11}, {0, 0, 2, 8}}, []int{1, 1}, []int{1, 4}, []int{2, 2}, []int{2, 4}}, []int{22, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubMatrixSumQueries(tt.args.A, tt.args.B, tt.args.C, tt.args.D, tt.args.E); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubMatrixSumQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubMatrixSumQueries-8   	 5658118	       213.8 ns/op	     168 B/op	       5 allocs/op
func BenchmarkSubMatrixSumQueries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubMatrixSumQueries([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 1}, []int{1, 2}, []int{2, 3}, []int{2, 3})
	}
}

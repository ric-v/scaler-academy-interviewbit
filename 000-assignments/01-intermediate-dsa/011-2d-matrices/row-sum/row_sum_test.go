package rowsum

import (
	"reflect"
	"testing"
)

func TestRowSum(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantSum []int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{6, 15, 24}},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, []int{6, 15}},
		{"test3", args{[][]int{{1, 2}, {4, 5}, {7, 8}}}, []int{3, 9, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := RowSum(tt.args.A); !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("RowSum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

// BenchmarkRowSum-8   	31526678	        34.04 ns/op	      24 B/op	       1 allocs/op
func BenchmarkRowSum(b *testing.B) {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		RowSum(A)
	}
}

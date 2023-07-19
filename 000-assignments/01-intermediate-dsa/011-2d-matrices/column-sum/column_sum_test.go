package columnsum

import (
	"reflect"
	"testing"
)

func TestColumnSum(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantSum []int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{12, 15, 18}},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, []int{5, 7, 9}},
		{"test3", args{[][]int{{1, 2}, {4, 5}, {7, 8}}}, []int{12, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := ColumnSum(tt.args.A); !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("ColumnSum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

// BenchmarkColumnSum-8   	30812446	        34.86 ns/op	      24 B/op	       1 allocs/op
func BenchmarkColumnSum(b *testing.B) {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		ColumnSum(A)
	}
}

package transposematrix

import (
	"reflect"
	"testing"
)

func TestTransposeMatrix(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{"test2", args{[][]int{{1, 2}, {1, 2}, {1, 2}}}, [][]int{{1, 1, 1}, {2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := TransposeMatrix(tt.args.A); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TransposeMatrix() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// BenchmarkTransposeMatrix-8   	 7582798	       156.9 ns/op	     152 B/op	       4 allocs/op
func BenchmarkTransposeMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TransposeMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}

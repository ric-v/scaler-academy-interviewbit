package matrixsubtraction

import (
	"reflect"
	"testing"
)

func TestMatrixSubtraction(t *testing.T) {
	type args struct {
		A [][]int
		B [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}}, [][]int{{-8, -6, -4}, {-2, 0, 2}, {4, 6, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MatrixSubtraction(tt.args.A, tt.args.B); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MatrixSubtraction() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// BenchmarkMatrixSubtraction-8   	 8074945	       148.5 ns/op	     152 B/op	       4 allocs/op
func BenchmarkMatrixSubtraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MatrixSubtraction([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}

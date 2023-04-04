package martrixscalarproduct

import (
	"reflect"
	"testing"
)

func TestMatrixScalarProduct(t *testing.T) {
	type args struct {
		A [][]int
		B int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2}, [][]int{{2, 4, 6}, {8, 10, 12}, {14, 16, 18}}},
		{"test2", args{[][]int{{1, 2}, {1, 2}, {1, 2}}, 3}, [][]int{{3, 6}, {3, 6}, {3, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MatrixScalarProduct(tt.args.A, tt.args.B); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MatrixScalarProduct() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// BenchmarkMatrixScalarProduct-8   	 8966116	       135.2 ns/op	     152 B/op	       4 allocs/op
func BenchmarkMatrixScalarProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MatrixScalarProduct([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2)
	}
}

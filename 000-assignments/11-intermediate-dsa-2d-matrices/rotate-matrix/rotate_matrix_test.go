package rotatematrix

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	type args struct {
		A *[][]int
	}
	tests := []struct {
		name string
		args args
		want *[][]int
	}{
		{"test1", args{&[][]int{{1, 2}, {3, 4}}}, &[][]int{{3, 1}, {4, 2}}},
		{"test2", args{&[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, &[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateMatrix(tt.args.A)
			if !reflect.DeepEqual(tt.args.A, tt.want) {
				t.Errorf("RotateMatrix() = %v, want %v", tt.args.A, tt.want)
			}
		})
	}
}

// BenchmarkRotateMatrix-8   	 7998111	       147.0 ns/op	     152 B/op	       4 allocs/op
func BenchmarkRotateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateMatrix(&[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}

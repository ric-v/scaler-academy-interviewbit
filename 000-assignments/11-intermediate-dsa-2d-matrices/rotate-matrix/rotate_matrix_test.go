package rotatematrix

import (
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
		{"test1", args{&[][]int{{3, 1}, {4, 2}}}, &[][]int{{1, 3}, {2, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateMatrix(tt.args.A)
			if !Equal(tt.args.A, tt.want) {
				t.Errorf("RotateMatrix() = %v, want %v", tt.args.A, tt.want)
			}
		})
	}
}

func Equal(a, b *[][]int) bool {
	if len(*a) != len(*b) {
		return false
	}
	for i := 0; i < len(*a); i++ {
		if len((*a)[i]) != len((*b)[i]) {
			return false
		}
		for j := 0; j < len((*a)[i]); j++ {
			if (*a)[i][j] != (*b)[i][j] {
				return false
			}
		}
	}
	return true
}

func BenchmarkRotateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateMatrix(&[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}

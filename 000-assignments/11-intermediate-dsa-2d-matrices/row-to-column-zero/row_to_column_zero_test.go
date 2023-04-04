package rowtocolumnzero

import (
	"reflect"
	"testing"
)

func TestRowToColumnZero(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}}, [][]int{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}}},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, [][]int{{1, 2, 3}, {4, 5, 6}}},
		{"test3", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 0}, {9, 2, 0, 4}}}, [][]int{{1, 2, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := RowToColumnZero(tt.args.A); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("RowToColumnZero() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// BenchmarkRowToColumnZero-8   	85771272	        13.42 ns/op	       0 B/op	       0 allocs/op
func BenchmarkRowToColumnZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RowToColumnZero([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}

package addthematrices

import (
	"reflect"
	"testing"
)

func TestAddTheMatrices(t *testing.T) {
	type args struct {
		A [][]int
		B [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, [][]int{{2, 4, 6}, {8, 10, 12}, {14, 16, 18}}},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}}, [][]int{{10, 10, 10}, {10, 10, 10}, {10, 10, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AddTheMatrices(tt.args.A, tt.args.B); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AddTheMatrices() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// BenchmarkAddTheMatrices-8   	 7193385	       178.0 ns/op	     152 B/op	       4 allocs/op
func BenchmarkAddTheMatrices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddTheMatrices([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}

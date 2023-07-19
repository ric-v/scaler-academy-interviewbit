package evennumbersinarange

import (
	"reflect"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	type args struct {
		A []int
		B [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{[]int{1, 2, 3, 4, 5}, [][]int{{0, 2}, {2, 4}, {1, 4}}}, []int{1, 1, 2}},
		{"test 2", args{[]int{2, 1, 8, 3, 9, 6}, [][]int{{0, 3}, {3, 5}, {1, 3}, {2, 4}}}, []int{2, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenNumbers(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkEvenNumbers-8   	11198239	       111.5 ns/op	      56 B/op	       3 allocs/op
func BenchmarkEvenNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenNumbers([]int{1, 2, 3, 4, 5}, [][]int{{0, 2}, {2, 4}, {1, 4}})
	}
}

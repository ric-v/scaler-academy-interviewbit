package multipleleftrotationsofthearray

import (
	"reflect"
	"testing"
)

func TestMultipleLeftRotationsOfTheArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, []int{2, 3}}, [][]int{{3, 4, 5, 1, 2}, {4, 5, 1, 2, 3}}},
		{"test2", args{[]int{5, 17, 100, 11}, []int{1}}, [][]int{{17, 100, 11, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultipleLeftRotationsOfTheArray(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultipleLeftRotationsOfTheArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMultipleLeftRotationsOfTheArray-8   	 9909632	       126.3 ns/op	     144 B/op	       3 allocs/op
func BenchmarkMultipleLeftRotationsOfTheArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultipleLeftRotationsOfTheArray([]int{1, 2, 3, 4, 5}, []int{2, 3})
	}
}

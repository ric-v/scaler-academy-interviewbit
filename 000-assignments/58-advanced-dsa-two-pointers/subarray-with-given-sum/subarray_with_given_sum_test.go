package subarraywithgivensum

import (
	"reflect"
	"testing"
)

func TestSubArrayWithGivenSum(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, 5}, []int{2, 3}},
		{"test2", args{[]int{5, 10, 20, 100, 105}, 110}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArrayWithGivenSum(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubArrayWithGivenSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubArrayWithGivenSum-8   	50493361	        23.67 ns/op	      16 B/op	       1 allocs/op
func BenchmarkSubArrayWithGivenSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubArrayWithGivenSum([]int{1, 2, 3, 4, 5}, 5)
	}
}

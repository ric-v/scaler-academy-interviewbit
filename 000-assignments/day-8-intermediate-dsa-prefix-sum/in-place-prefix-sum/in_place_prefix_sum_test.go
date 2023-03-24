package inplaceprefixsum

import (
	"reflect"
	"testing"
)

func TestInPlacePrefixSum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}}, []int{1, 3, 6, 10, 15}},
		{"test2", args{[]int{4, 3, 2}}, []int{4, 7, 9}},
		{"test_large", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}, []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 66, 78, 91, 105, 120, 136, 153, 171, 190, 210}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InPlacePrefixSum(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InPlacePrefixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkInPlacePrefixSum-8   	35161125	        30.43 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInPlacePrefixSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InPlacePrefixSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	}
}

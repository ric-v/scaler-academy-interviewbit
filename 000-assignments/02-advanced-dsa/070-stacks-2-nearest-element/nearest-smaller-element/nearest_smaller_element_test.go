package nearestsmallerelement

import (
	"reflect"
	"testing"
)

func TestNearestSmallerElement(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{4, 5, 2, 10, 8}}, []int{-1, 4, -1, 2, 2}},
		{"test2", args{[]int{3, 2, 1}}, []int{-1, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NearestSmallerElement(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NearestSmallerElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkNearestSmallerElement-8   	18122014	        67.95 ns/op	      96 B/op	       2 allocs/op
func BenchmarkNearestSmallerElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NearestSmallerElement([]int{4, 5, 2, 10, 8})
	}
}

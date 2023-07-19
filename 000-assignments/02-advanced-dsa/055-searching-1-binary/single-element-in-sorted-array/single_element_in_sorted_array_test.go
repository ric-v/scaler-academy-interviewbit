package singleelementinsortedarray

import "testing"

func TestSingleElementInSortedArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 2, 2, 3, 3, 4, 4, 8, 8, 9}}, 9},
		{"test2", args{[]int{3, 3, 7, 7, 10, 11, 11}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleElementInSortedArray(tt.args.A); got != tt.want {
				t.Errorf("SingleElementInSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSingleElementInSortedArray-8   	191057094	         6.262 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSingleElementInSortedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleElementInSortedArray([]int{1, 1, 2, 2, 3, 3, 4, 4, 8, 8, 9})
	}
}

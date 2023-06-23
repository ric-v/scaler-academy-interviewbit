package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr   []int
		query int
	}
	tests := []struct {
		name        string
		args        args
		wantElement int
		wantIndex   int
	}{
		{"test1", args{[]int{1, 3, 5, 6}, 5}, 5, 2},
		{"test2", args{[]int{1, 4, 9}, 3}, 3, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotElement, gotIndex := BinarySearch(tt.args.arr, tt.args.query)
			if gotElement != tt.wantElement {
				t.Errorf("BinarySearch() gotElement = %v, want %v", gotElement, tt.wantElement)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("BinarySearch() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

// BenchmarkBinarySearch-8   	503058385	         2.457 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch([]int{1, 3, 5, 6}, 5)
	}
}

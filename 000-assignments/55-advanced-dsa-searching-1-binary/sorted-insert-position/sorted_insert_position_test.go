package sortedinsertposition

import "testing"

func TestSortedInsertPosition(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"test2", args{[]int{1, 4, 9}, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedInsertPosition(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("SortedInsertPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSortedInsertPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortedInsertPosition([]int{1, 3, 5, 6}, 5)
	}
}

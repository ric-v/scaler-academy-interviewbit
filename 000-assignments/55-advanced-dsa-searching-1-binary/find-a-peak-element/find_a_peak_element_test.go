package findapeakelement

import "testing"

func TestFindAPeakElement(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}}, 5},
		{"test2", args{[]int{5, 17, 100, 11}}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAPeakElement(tt.args.A); got != tt.want {
				t.Errorf("FindAPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFindAPeakElement-8   	387023613	         3.127 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFindAPeakElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindAPeakElement([]int{1, 2, 3, 4, 5})
	}
}

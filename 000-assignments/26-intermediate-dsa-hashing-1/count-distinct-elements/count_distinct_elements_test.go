package countdistinctelements

import "testing"

func TestCountDistinctElements(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{3, 4, 3, 6, 6}}, 3},
		{"test2", args{[]int{3, 3, 3, 9, 0, 1, 0}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDistinctElements(tt.args.A); got != tt.want {
				t.Errorf("CountDistinctElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCountDistinctElements-8   	19753672	        57.99 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountDistinctElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountDistinctElements([]int{3, 4, 3, 6, 6})
	}
}

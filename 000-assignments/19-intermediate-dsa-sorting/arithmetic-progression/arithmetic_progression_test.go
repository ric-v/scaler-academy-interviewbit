package arithmeticprogression

import "testing"

func TestArithmeticProgression(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{3, 5, 1}}, 1},
		{"test2", args{[]int{2, 4, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArithmeticProgression(tt.args.A); got != tt.want {
				t.Errorf("ArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkArithmeticProgression-8   	12430863	       100.3 ns/op	      48 B/op	       2 allocs/op
func BenchmarkArithmeticProgression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArithmeticProgression([]int{3, 5, 1})
	}
}

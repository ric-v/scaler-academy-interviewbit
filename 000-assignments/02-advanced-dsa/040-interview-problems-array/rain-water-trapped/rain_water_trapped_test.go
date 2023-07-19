package rainwatertrapped

import "testing"

func TestRainWaterTrapped(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 2}}, 1},
		{"test2", args{[]int{1, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RainWaterTrapped(tt.args.A); got != tt.want {
				t.Errorf("RainWaterTrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRainWaterTrapped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RainWaterTrapped([]int{0, 1, 0, 2})
	}
}

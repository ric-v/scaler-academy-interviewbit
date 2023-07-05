package aggressivecows

import "testing"

func TestAggresssiveCowns(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, 3}, 2},
		{"test2", args{[]int{1, 2}, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AggresssiveCowns(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("AggresssiveCowns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAggresssiveCowns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AggresssiveCowns([]int{1, 2, 3, 4, 5}, 3)
	}
}

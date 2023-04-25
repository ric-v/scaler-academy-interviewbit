package divisibilityby3

import "testing"

func TestDivisibilityBy3(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3}}, 1},
		{"test2", args{[]int{1, 0, 0, 1, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivisibilityBy3(tt.args.A); got != tt.want {
				t.Errorf("DivisibilityBy3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDivisibilityBy3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivisibilityBy3([]int{1, 0, 0, 1, 2})
	}
}

package maxandmin

import "testing"

func TestMaxAndMin(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1}}, 0},
		{"test2", args{[]int{4, 7, 3, 8}}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAndMin(tt.args.A); got != tt.want {
				t.Errorf("MaxAndMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

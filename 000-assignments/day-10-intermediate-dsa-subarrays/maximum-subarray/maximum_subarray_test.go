package maximumsubarray

import "testing"

func TestMaximumSubArray(t *testing.T) {
	type args struct {
		A int
		B int
		C []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5, 12, []int{2, 1, 3, 4, 5}}, 12},
		{"test2", args{3, 1, []int{2, 2, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubArray(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("MaximumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

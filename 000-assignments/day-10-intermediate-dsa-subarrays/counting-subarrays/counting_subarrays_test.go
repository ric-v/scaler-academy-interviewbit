package countingsubarrays

import "testing"

func TestCountingSubarrays(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 5, 6}, 10}, 4},
		{"test2", args{[]int{1, 11, 2, 3, 15}, 10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountingSubarrays(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("CountingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

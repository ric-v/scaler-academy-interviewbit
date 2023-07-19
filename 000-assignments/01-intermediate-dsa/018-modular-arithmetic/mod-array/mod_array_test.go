package modarray

import "testing"

func TestModArray(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 4, 3}, 2}, 1},
		{"test2", args{[]int{4, 3, 5, 3, 5, 3, 2, 1}, 47}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModArray(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("ModArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

package pickfrombothsides

import "testing"

func TestPickFromBothSides(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{5, -2, 3, 1, 2}, 3}, 8},
		{"test2", args{[]int{2, 3, -1, 4, 2, 1}, 4}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickFromBothSides(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("PickFromBothSides() = %v, want %v", got, tt.want)
			}
		})
	}
}

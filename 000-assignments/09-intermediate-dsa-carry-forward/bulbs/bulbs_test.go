package bulbs

import "testing"

func TestBulbs(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 1}}, 3},
		// {"test2", args{[]int{1, 1, 1, 1}}, 0},
		// {"test3", args{[]int{0, 0, 0, 0}}, 1},
		// {"test4", args{[]int{0, 1, 0, 1, 0, 1, 0, 1}}, 8},
		// {"test5", args{[]int{0, 1, 0, 1, 0, 1, 0, 1, 0}}, 9},
		// {"test6", args{[]int{0, 0, 0, 1}}, 2},
		// {"test_large", args{[]int{1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1}}, 70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bulbs(tt.args.A); got != tt.want {
				t.Errorf("Bulbs() = %v, want %v", got, tt.want)
			}
		})
	}
}

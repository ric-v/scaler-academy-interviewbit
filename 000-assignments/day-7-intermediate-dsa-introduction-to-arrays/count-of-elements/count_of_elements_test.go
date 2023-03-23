package countofelements

import "testing"

func TestCountOfElements(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test2", args{[]int{3, 1, 2}}, 2},
		{"test2", args{[]int{5, 5, 3}}, 1},
		{"test3", args{[]int{1, 1, 1}}, 0},
		{"test4", args{[]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6}}, 8},
		{"test4", args{[]int{1, 2, 2, 3, 8, 4, 5, 7, 6}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOfElements(tt.args.A); got != tt.want {
				t.Errorf("CountOfElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

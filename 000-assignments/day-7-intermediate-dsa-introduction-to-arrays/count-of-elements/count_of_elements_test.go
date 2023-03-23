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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOfElements(tt.args.A); got != tt.want {
				t.Errorf("CountOfElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

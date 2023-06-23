package sumthedifference

import "testing"

func TestSumTheDifference(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2}}, 1},
		{"test2", args{[]int{3, 5, 10}}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumTheDifference(tt.args.A); got != tt.want {
				t.Errorf("SumTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

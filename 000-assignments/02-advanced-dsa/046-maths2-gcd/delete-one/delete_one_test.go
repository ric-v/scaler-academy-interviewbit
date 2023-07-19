package deleteone

import "testing"

func TestDeleteOne(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{12, 15, 18}}, 6},
		{"test2", args{[]int{5, 15, 30}}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteOne(tt.args.A); got != tt.want {
				t.Errorf("DeleteOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

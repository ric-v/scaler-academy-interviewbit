package evensubarrays

import "testing"

func TestEvenSubarrays(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{[]int{2, 4, 8, 6}}, "YES"},
		{"test2", args{[]int{2, 4, 8, 7, 6}}, "NO"},
		{"test2", args{[]int{2, 5, 8, 4, 7, 6}}, "YES"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenSubarrays(tt.args.A); got != tt.want {
				t.Errorf("EvenSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

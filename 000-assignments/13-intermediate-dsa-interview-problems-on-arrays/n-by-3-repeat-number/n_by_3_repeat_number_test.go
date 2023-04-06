package nby3repeatnumber

import "testing"

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 1, 1}}, 1},
		{"test2", args{[]int{1, 2, 3}}, -1},
		{"test3", args{[]int{1, 1, 1, 2, 3, 5, 7}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMajorityElement-8   	100000000	        11.82 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMajorityElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MajorityElement([]int{1, 2, 3, 1, 1})
	}
}

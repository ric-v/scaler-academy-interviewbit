package missingnumber

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	type args struct {
		in0 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{3, 1}}, 2},
		{"test2", args{[]int{9, 6, 4, 2, 3, 5, 7, 1}}, 8},
		{"test1", args{[]int{2, 3, 1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.args.in0); got != tt.want {
				t.Errorf("MissingNumber() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}

// 3951795               313.4 ns/op           104 B/op          2 allocs/op
func BenchmarkMissingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	}
}

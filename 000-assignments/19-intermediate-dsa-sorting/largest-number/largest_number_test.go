package largestnumber

import "testing"

func TestLargestNumber(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1 ", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
		{"test2 ", args{[]int{2, 3, 9, 0}}, "9320"},
		{"test3 ", args{[]int{0, 0, 0, 0}}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestNumber(tt.args.A); got != tt.want {
				t.Errorf("LargestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLargestNumber-8   	 1645160	       693.0 ns/op	     144 B/op	       4 allocs/op
func BenchmarkLargestNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestNumber([]int{3, 30, 34, 5, 9})
	}
}

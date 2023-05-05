package findfibonacciii

import "testing"

func TestFindFibonacci(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{0}, 0},
		{"test2", args{1}, 1},
		{"test3", args{2}, 1},
		{"test4", args{3}, 2},
		{"test5", args{4}, 3},
		{"test6", args{5}, 5},
		{"test7", args{6}, 8},
		{"test8", args{7}, 13},
		{"test9", args{8}, 21},
		{"test10", args{9}, 34},
		{"test11", args{10}, 55},
		{"test12", args{20}, 6765},
		{"test15", args{30}, 832040},
		{"test16", args{40}, 102334155},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFibonacci(tt.args.A); got != tt.want {
				t.Errorf("FindFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFindFibonacci-8   	       3	 490419188 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFindFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindFibonacci(40)
	}
}

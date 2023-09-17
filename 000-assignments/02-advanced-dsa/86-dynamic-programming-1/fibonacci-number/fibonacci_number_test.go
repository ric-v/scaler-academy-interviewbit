package fibonaccinumber

import "testing"

func TestFibonacciNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2}, 1},
		{"test2", args{3}, 2},
		{"test3", args{4}, 3},
		{"test4", args{5}, 5},
		{"test5", args{6}, 8},
		{"test6", args{7}, 13},
		{"test7", args{8}, 21},
		{"test8", args{9}, 34},
		{"test9", args{10}, 55},
		{"test10", args{11}, 89},
		{"test11", args{12}, 144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciNumber(tt.args.n); got != tt.want {
				t.Errorf("FibonacciNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFibonacciNumber-8   	14736980	        84.87 ns/op	      96 B/op	       1 allocs/op
func BenchmarkFibonacciNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciNumber(10)
	}
}

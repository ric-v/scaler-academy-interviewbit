package countofprime

import "testing"

func TestCountOfPrime(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Case 1", args{19}, 8},
		{"Test Case 2", args{1}, 0},
		{"Test Case 3", args{2}, 1},
		{"Test Case 4", args{3}, 2},
		{"Test Case 5", args{4}, 2},
		{"Test Case 6", args{5}, 3},
		{"Test Case 7", args{6}, 3},
		{"Test Case 8", args{7}, 4},
		{"Test Case 9", args{8}, 4},
		{"Test Case 10", args{9}, 4},
		{"Test Case 11", args{10}, 4},
		{"Test Case 12", args{100}, 25},
		{"Test Case 13", args{1000}, 168},
		{"Test Case 14", args{10000}, 1229},
		{"Test Case 15", args{100000}, 9592},
		{"Test Case 16", args{1000000}, 78498},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOfPrime(tt.args.A); got != tt.want {
				t.Errorf("CountOfPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCountOfPrime-8   	       1	4973788566 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountOfPrime_1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountOfPrime(1000000)
	}
}

// BenchmarkCountOfPrime_1000-8        6404            189136 ns/op               0 B/op          0 allocs/op
func BenchmarkCountOfPrime_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountOfPrime(1000)
	}
}

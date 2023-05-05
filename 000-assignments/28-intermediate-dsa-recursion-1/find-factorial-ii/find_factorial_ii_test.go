package findfactorialii

import "testing"

func TestFindFactorial(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1}, 1},
		{"test2", args{2}, 2},
		{"test3", args{3}, 6},
		{"test4", args{4}, 24},
		{"test5", args{5}, 120},
		{"test6", args{6}, 720},
		{"test7", args{7}, 5040},
		{"test8", args{8}, 40320},
		{"test9", args{9}, 362880},
		{"test10", args{10}, 3628800},
		{"test11", args{20}, 2432902008176640000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFactorial(tt.args.A); got != tt.want {
				t.Errorf("FindFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFindFactorial-8   	70694876	        16.81 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFindFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindFactorial(10)
	}
}

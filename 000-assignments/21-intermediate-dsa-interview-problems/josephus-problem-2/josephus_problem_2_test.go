package josephusproblem2

import "testing"

func TestJosephusProblem2(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{4}, 1},
		{"test2", args{5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JosephusProblem2(tt.args.A); got != tt.want {
				t.Errorf("JosephusProblem2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkJosephusProblem2-8   	34249701	        33.89 ns/op	       0 B/op	       0 allocs/op
func BenchmarkJosephusProblem2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JosephusProblem2(100)
	}
}

package perfectnumbers

import "testing"

func TestPerfectNumbers(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{2}, "22"},
		{"test2", args{3}, "1111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PerfectNumbers(tt.args.A); got != tt.want {
				t.Errorf("PerfectNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPerfectNumbers-8   	 4304654	       282.6 ns/op	     124 B/op	       7 allocs/op
func BenchmarkPerfectNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PerfectNumbers(2)
	}
}

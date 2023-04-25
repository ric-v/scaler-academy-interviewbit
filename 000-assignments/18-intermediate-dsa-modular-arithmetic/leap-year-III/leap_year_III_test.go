package leapyeariii

import "testing"

func TestLeapYearIII(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2016}, 1},
		{"test2", args{2017}, 0},
		{"test3", args{2200}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeapYearIII(tt.args.A); got != tt.want {
				t.Errorf("LeapYearIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLeapYearIII-8   	1000000000	         0.3308 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLeapYearIII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeapYearIII(2016)
	}
}

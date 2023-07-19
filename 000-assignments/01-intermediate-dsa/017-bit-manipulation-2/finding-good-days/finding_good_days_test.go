package findinggooddays

import "testing"

func TestFindingGoodDays(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5}, 2},
		{"test2", args{8}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindingGoodDays(tt.args.A); got != tt.want {
				t.Errorf("FindingGoodDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFindingGoodDays-8   	61778173	        18.61 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFindingGoodDays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindingGoodDays(100000000)
	}
}

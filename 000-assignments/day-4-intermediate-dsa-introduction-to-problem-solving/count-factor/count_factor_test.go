package count

import (
	"testing"
)

func TestCountFactor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		{"1", args{1}, 1},
		{"2", args{2}, 2},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"5", args{5}, 2},
		{"10", args{10}, 4},
		{"100", args{100}, 9},
		{"1000", args{1000}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := CountFactor(tt.args.n); gotC != tt.wantC {
				t.Errorf("CountFactor() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

// BenchmarkCountFactor-8             63378             18544 ns/op               0 B/op          0 allocs/op
func BenchmarkCountFactor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountFactor(1000000000)
	}
}

func TestBruteForce(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		{"1", args{1}, 1},
		{"2", args{2}, 2},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"5", args{5}, 2},
		{"10", args{10}, 4},
		{"100", args{100}, 9},
		{"1000", args{1000}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := BruteForce(tt.args.n); gotC != tt.wantC {
				t.Errorf("BruteForce() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

// BenchmarkBruteForce-8                  4         287210445 ns/op               0 B/op          0 allocs/op
func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(1000000000)
	}
}

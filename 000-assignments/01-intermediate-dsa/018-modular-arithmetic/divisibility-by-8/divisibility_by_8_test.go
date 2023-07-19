package divisibilityby8

import "testing"

func TestDivisibilityBy8(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"16"}, 1},
		{"test2", args{"123"}, 0},
		{"test3", args{"611238745789"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivisibilityBy8(tt.args.A); got != tt.want {
				t.Errorf("DivisibilityBy8() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkDivisibilityBy8-8   	190330342	         6.410 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDivisibilityBy8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivisibilityBy8("611238745789")
	}
}

package specialsubsequencesag2

import "testing"

func TestSpecialSubsequencesAG2(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{"ABCGAG"}, 3},
		{"test2", args{"GAB"}, 0},
		{"test3", args{"AGAGAGAG"}, 10},
		{"test4_large", args{"ABCGDEFGHIJKAGLMNOPQRSTUVWXYZAG"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpecialSubsequencesAG2(tt.args.A); got != tt.want {
				t.Errorf("SpecialSubsequencesAG2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSpecialSubsequencesAG2-8   	65000839	        19.17 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSpecialSubsequencesAG2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SpecialSubsequencesAG2("ABCGDEFGHIJKAGLMNOPQRSTUVWXYZAG")
	}
}

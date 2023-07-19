package helpfromsam

import "testing"

func TestHelpFromSam(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{3}, 2},
		{"test2", args{6}, 2},
		{"test3", args{7}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelpFromSam(tt.args.A); got != tt.want {
				t.Errorf("HelpFromSam() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkHelpFromSam-8   	32299479	        35.42 ns/op	       0 B/op	       0 allocs/op
func BenchmarkHelpFromSam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelpFromSam(100000000)
	}
}

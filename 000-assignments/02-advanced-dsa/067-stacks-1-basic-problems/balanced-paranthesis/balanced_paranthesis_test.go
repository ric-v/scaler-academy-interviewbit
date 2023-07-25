package balancedparanthesis

import "testing"

func TestBalancedParanthesis(t *testing.T) {
	type args struct {
		A string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{"()"}, 1},
		{"1", args{"(())"}, 1},
		{"2", args{"()()"}, 1},
		{"3", args{"{[[}[]}(]}){"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BalancedParanthesis(tt.args.A); got != tt.want {
				t.Errorf("BalancedParanthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkBalancedParanthesis-8   	18202070	        62.29 ns/op	      24 B/op	       2 allocs/op
func BenchmarkBalancedParanthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BalancedParanthesis("{[[}[]}(]}){")
	}
}

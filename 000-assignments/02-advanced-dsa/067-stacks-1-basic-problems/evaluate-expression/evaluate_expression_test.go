package evaluateexpression

import "testing"

func TestEvaluateExpression(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]string{"2", "1", "+", "3", "*"}}, 9},
		{"test2", args{[]string{"4", "13", "5", "/", "+"}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluateExpression(tt.args.A); got != tt.want {
				t.Errorf("EvaluateExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkEvaluateExpression-8   	13749830	        86.15 ns/op	      24 B/op	       2 allocs/op
func BenchmarkEvaluateExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvaluateExpression([]string{"2", "1", "+", "3", "*"})
	}
}

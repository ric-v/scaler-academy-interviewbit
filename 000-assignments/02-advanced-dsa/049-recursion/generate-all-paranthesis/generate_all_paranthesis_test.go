package generateallparanthesis

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{1}, []string{"()"}},
		{"test2", args{2}, []string{"(())", "()()"}},
		{"test3", args{3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_generateParenthesis-8   	 1388409	       884.8 ns/op	     352 B/op	      24 allocs/op
func Benchmark_generateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(3)
	}
}

package palindromelist

import (
	ll "linkedlist"
	"testing"
)

func TestPalindromeList(t *testing.T) {
	type args struct {
		A *ll.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{ll.NewNode(1, 2, 2, 1)}, 1},
		{"test2", args{ll.NewNode(1, 3, 2)}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromeList(tt.args.A); got != tt.want {
				t.Errorf("PalindromeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPalindromeList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeList(ll.NewNode(1, 2, 2, 1))
	}
}

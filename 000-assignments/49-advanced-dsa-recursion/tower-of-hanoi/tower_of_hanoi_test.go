package towerofhanoi

import (
	"reflect"
	"testing"
)

func Test_towerOfHanoi(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{2}, [][]int{{1, 1, 2}, {2, 1, 3}, {1, 2, 3}}},
		{"test2", args{3}, [][]int{{1, 1, 3}, {2, 1, 2}, {1, 3, 2}, {3, 1, 3}, {1, 2, 1}, {2, 2, 3}, {1, 1, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := towerOfHanoi(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("towerOfHanoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkTowerOfHanoi-8   	 4162908	       293.3 ns/op	     344 B/op	       8 allocs/op
func BenchmarkTowerOfHanoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		towerOfHanoi(3)
	}
}

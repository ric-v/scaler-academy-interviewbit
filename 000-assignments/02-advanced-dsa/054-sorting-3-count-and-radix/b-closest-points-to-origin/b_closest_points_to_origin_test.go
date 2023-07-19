package bclosestpointstoorigin

import (
	"reflect"
	"testing"
)

func TestClosestPointsToOrigin(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[][]int{{1, 3}, {-2, 2}}, 1}, [][]int{{-2, 2}}},
		{"test2", args{[][]int{{1, -1}, {2, -1}}, 1}, [][]int{{1, -1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClosestPointsToOrigin(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClosestPointsToOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkClosestPointsToOrigin-8         2335642               516.3 ns/op           264 B/op         11 allocs/op
func BenchmarkClosestPointsToOrigin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClosestPointsToOrigin([][]int{{1, 3}, {-2, 2}}, 1)
	}
}

package mergeoverlappingintervals

import (
	"reflect"
	"testing"
)

func TestMergeOverlappingIntervals(t *testing.T) {
	type args struct {
		A []Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"test1", args{[]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, []Interval{{1, 6}, {8, 10}, {15, 18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeOverlappingIntervals(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeOverlappingIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

//  4157031               291.8 ns/op           264 B/op          7 allocs/op
func BenchmarkMergeOverlappingIntervals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeOverlappingIntervals([]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	}
}

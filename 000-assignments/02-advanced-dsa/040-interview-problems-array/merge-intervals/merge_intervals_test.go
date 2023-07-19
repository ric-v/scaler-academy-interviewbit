package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	type args struct {
		intervals   []Interval
		newInterval Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{"test1", args{[]Interval{{1, 3}, {6, 9}}, Interval{2, 5}}, []Interval{{1, 5}, {6, 9}}},
		{"test2", args{[]Interval{{1, 3}, {6, 9}}, Interval{2, 6}}, []Interval{{1, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntervals(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 18810958                65.75 ns/op           48 B/op          2 allocs/op
func BenchmarkMergeIntervals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeIntervals([]Interval{{1, 3}, {6, 9}}, Interval{2, 5})
	}
}

package addonetonumber

import (
	"reflect"
	"testing"
)

func TestAddOneToNumber(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"test2", args{[]int{9, 9, 9}}, []int{1, 0, 0, 0}},
		{"test4", args{[]int{6, 4, 7, 7, 8, 9}}, []int{6, 4, 7, 7, 9, 0}},
		{"test5", args{[]int{0, 3, 7, 6, 4, 0, 5, 5, 5}}, []int{3, 7, 6, 4, 0, 5, 5, 6}},
		{"test6", args{[]int{9, 8, 6, 0, 3, 3, 7, 4, 3, 7, 9}}, []int{9, 8, 6, 0, 3, 3, 7, 4, 3, 8, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddOneToNumber(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOneToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddOneToNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddOneToNumber([]int{1, 2, 3})
	}
}

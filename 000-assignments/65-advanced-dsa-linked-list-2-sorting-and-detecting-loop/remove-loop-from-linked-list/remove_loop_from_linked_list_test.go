package removeloopfromlinkedlist

import (
	"reflect"
	"testing"
)

func TestRemoveLoopFromLinkedList(t *testing.T) {
	type args struct {
		A *ListNode
	}

	test1 := ListNode_new(1)
	test1.next = ListNode_new(2)
	want1 := ListNode_new(1)
	want1.next = ListNode_new(2)
	test1.next.next = test1

	test2 := ListNode_new(3)
	test2.next = ListNode_new(2)
	test2.next.next = ListNode_new(4)
	test2.next.next.next = ListNode_new(5)
	test2.next.next.next.next = ListNode_new(6)
	test2.next.next.next.next.next = test2.next.next.next
	want2 := ListNode_new(3)
	want2.next = ListNode_new(2)
	want2.next.next = ListNode_new(4)
	want2.next.next.next = ListNode_new(5)
	want2.next.next.next.next = ListNode_new(6)

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test1", args{test1}, want1},
		{"test2", args{test2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveLoopFromLinkedList(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLoopFromLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRemoveLoopFromLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

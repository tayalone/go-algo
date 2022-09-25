package twopointers

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// // // // // //
	// l2 := &ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }
	// // // // // //
	// // // // // //

	// // // // // //
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "Case 2",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				n: 0,
			},
			want: nil,
		},
		{
			name: "Case 1",
			args: args{
				head: nil,
				n:    0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

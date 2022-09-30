package linkedlist

import (
	"reflect"
	"testing"
)

func makeWant(got *ListNode) []int {
	r := []int{}
	for got != nil {
		r = append(r, got.Val)
		got = got.Next
	}

	return r
}

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: []int{},
		},
		{
			name: "Case 2",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp := MergeTwoLists(tt.args.list1, tt.args.list2)

			if got := makeWant(tmp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

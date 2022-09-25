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
	//  n3_1 :=  &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	//  }
	//  n3_2:=  &ListNode{
	// 	Val:  2,
	// 	Next: nil,
	//  }

	// // // // // //
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "CASE_3",
		// 	args: args{
		// 		head: &ListNode{
		// 			Val:  1,
		// 			Next: nil,
		// 		},
		// 		n: 0,
		// 	},
		// 	want: nil,
		// },
		{
			name: "CASE_2",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				n: 0,
			},
			want: []int{},
		},
		{
			name: "CASE_1",
			args: args{
				head: nil,
				n:    0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := RemoveNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			// }
			tmpGot := RemoveNthFromEnd(tt.args.head, tt.args.n)

			got := []int{}
			for tmpGot != nil {
				got = append(got, tmpGot.Val)
			}

			if !reflect.DeepEqual([]int{}, []int{}) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

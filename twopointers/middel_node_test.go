package twopointers

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	// // ---------- case 1 ------------------------------
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	want1 := head1.Next.Next
	// // ------------------------------------------------
	// // ------------ case 2 ----------------------------
	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	want2 := head2.Next.Next.Next
	// // -------------------------------------------------
	// // ------------ case 3 -----------------------------
	head3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	want3 := head3.Next
	// // --------------------------------------------------
	// // ------------ case 3 -----------------------------
	head4 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	want4 := head4.Next
	// // --------------------------------------------------
	head5 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	want5 := head5.Next.Next

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "Case 0",
			args: args{
				head: head5,
			},
			want: want5,
		},
		{
			name: "Case 1",
			args: args{
				head: head1,
			},
			want: want1,
		},
		{
			name: "Case 2",
			args: args{
				head: head2,
			},
			want: want2,
		}, {
			name: "Case 3",
			args: args{
				head: head3,
			},
			want: want3,
		}, {
			name: "Case 4",
			args: args{
				head: head4,
			},
			want: want4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

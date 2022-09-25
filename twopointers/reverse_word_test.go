package twopointers

import "testing"

//"Let's take LeetCode contest"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Case: 1", args: args{s: "Let's"}, want: "s'teL"},
		{name: "Case: 2", args: args{s: "take"}, want: "ekat"},
		{name: "Case: 3", args: args{s: "LeetCode"}, want: "edoCteeL"},
		{name: "Case: 3", args: args{s: "contest"}, want: "tsetnoc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Case: 1", args: args{s: "Let's take LeetCode contest"}, want: "s'teL ekat edoCteeL tsetnoc"},
		{name: "Case: 2", args: args{s: "God Ding"}, want: "doG gniD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.args.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

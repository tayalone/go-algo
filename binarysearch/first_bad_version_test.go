package binarysearch

import (
	"reflect"
	"testing"

	"github.com/tayalone/go-algo/binarysearch/mocks"
)

func TestNewFBV(t *testing.T) {
	type args struct {
		ibv mocks.IBV
	}
	tests := []struct {
		name string
		args args
		want FBVStruct
	}{
		// TODO: Add test cases.
		{
			name: "Case: New FBV",
			args: args{
				ibv: mocks.NewIbv(4),
			},
			want: FBVStruct{ibv: mocks.NewIbv(4)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFBV(tt.args.ibv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFBV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFBVStruct_FirstBadVersion(t *testing.T) {
	type fields struct {
		ibv mocks.IBV
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "Case: FirstBadVersion 1",
			fields: fields{
				ibv: mocks.NewIbv(100),
			},
			args: args{
				n: 1000,
			},
			want: 100,
		}, {
			name: "Case: FirstBadVersion 2",
			fields: fields{
				ibv: mocks.NewIbv(650),
			},
			args: args{
				n: 1000,
			},
			want: 650,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FBVStruct{
				ibv: tt.fields.ibv,
			}
			if got := f.FirstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("FBVStruct.FirstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

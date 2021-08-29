package slice

import (
	"reflect"
	"testing"
)

func TestAppendPos(t *testing.T) {
	type args struct {
		src  []interface{}
		item interface{}
		pos  int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "test1",
			args: args{
				src: []interface{}{
					1, 2, 3,
				},
				item: 100,
				pos:  1,
			},
			want: []interface{}{
				1, 100, 2, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendPos(tt.args.src, tt.args.item, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				got[0] = 10000
				t.Errorf("AppendPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveTail(t *testing.T) {
	type args struct {
		src   []interface{}
		index int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "test1",
			args: args{
				src:   []interface{}{1, 2, 3},
				index: 1,
			},
			want: []interface{}{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveTail(tt.args.src, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveTail() = %v, want %v", got, tt.want)
			}
		})
	}
}

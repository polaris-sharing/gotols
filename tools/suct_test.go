package tools

import (
	"reflect"
	"testing"
)

func TestGetStructByName(t *testing.T) {
	type args struct {
		su         interface{}
		structName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test1",
			args: args{
				su: struct {
					Name string
					Age  int
				}{
					"lucky",
					18,
				},
				structName: "Name",
			},
			want: "lucky",
		},
		{
			name: "test2",
			args: args{
				su: struct {
					Name string
					Age  int
				}{
					"lucky",
					18,
				},
				structName: "Age",
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStructByName(tt.args.su, tt.args.structName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStructByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

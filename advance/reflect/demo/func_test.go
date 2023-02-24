package demo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sjxiang/sharp/advance/reflect/types"
)

func TestIterateFuncs(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*FuncInfo
		wantErr error
	}{
		{
			name: "nil",
			wantErr: errors.New("输入 nil"),
		},  
		{
			name: "basic types",
			args: args{
				val: 123,
			},
			wantErr: errors.New("不支持类型"),
		},
		{
			name: "struct type",
			args: args{
				val: types.UserInfo{},
			},
			want: map[string]*FuncInfo{

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFuncs(tt.args.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
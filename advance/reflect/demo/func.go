package demo

import (
	"errors"
	"reflect"
)

// 输出方法信息，并执行调用
/*
	考虑输入可能是 nil、基础类型 ... （不支持）
	
		nil 
	
		结构体指针


*/
func IterateFuncs(input interface{}) (map[string]*FuncInfo, error) {
	if input == nil {
		return nil, errors.New("输入 nil")
	}

	typ := reflect.TypeOf(input)
	
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("不支持类型")
	}


	return nil, nil
}


type FuncInfo struct {
	Name   string
	In     []reflect.Type
	Out    []reflect.Type

	// 反射调用得到的结果
	Result []interface{}
}
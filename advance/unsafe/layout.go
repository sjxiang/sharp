package main

import (
	"fmt"
	"reflect"
)


func main() {
	e := Example {
		Name: "布局",
	}
	
	PrintFieldOffset(e)
}



// PrintFieldOffset 打印字段偏移量
// 用于研究内存布局
// 只接受结构体输入
func PrintFieldOffset(entity interface{}) {	
	
	typ := reflect.TypeOf(entity)

	if typ.Kind() != reflect.Struct {
		panic("非法类型")
	}

	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)
		fmt.Printf("%s - %d \n", fd.Name, fd.Offset)
	}
}


type Example struct {
	Name   string
	Price  float64      
}


## sharp，锋利


```

```

```
web
orm
rpc

微服务
缓存

```

Context 任务

并发

反射

unsafe

Go 测试 -- benchmark 测试 

	- benchmark 测试入门
	- 内存和 CPU 分析
	- Go 内存逃逸分析


网络编程
	- 网络编程入门
	- 连接池设计

SQL 
	database/sql
	gorm
	
AST
	规则引擎
模板

    wire 依赖注入


刷 Kpi，是设计解决通用方案，而不是一个个需求做过去，麻了。


# 脑洞 1

DTO 转 PO，不大实际，也没多少意义

	请求进来，需要对参数校验 （甚至搞点勾子或者前端处理，e.g. 头像变成 oss 链接，塞入 MySQL 存储）
	响应返回，需要序列化包装下（过滤）


# 脑洞 2







34:00


package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(val any) {
	// 复杂逻辑
	res, err := iterateFields(val)

	// 简单逻辑
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}

// iterateFields 返回所有的字段名字
// val 只能是结构体，或者结构体指针，可以是多重指针
func iterateFields(input any) (map[string]any, error) {
	typ := reflect.TypeOf(input)
	val := reflect.ValueOf(input)

	// 处理指针，要拿到指针指向的东西
	// 这里我们综合考虑了多重指针的效果
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	// 如果不是结构体，就返回 error
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("非法类型")
	}

	num := typ.NumField()
	res := make(map[string]any, num)
	for i := 0; i < num; i++ {
		fd := typ.Field(i)
		fdVal := val.Field(i)
		if fd.IsExported() {
			res[fd.Name] = fdVal.Interface()
		} else {
			// 为了演示效果，不公开字段我们用零值来填充
			res[fd.Name] = reflect.Zero(fd.Type).Interface()
		}
	}
	return res, nil
}

func SetField(entity any, field string, newVal any) error {
	val := reflect.ValueOf(entity)
	typ := val.Type()

	// 只能是一级指针，类似 *User
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("非法类型")
	}

	typ = typ.Elem()
	val = val.Elem()

	// 这个地方判断不出来 field 在不在
	fd := val.FieldByName(field)
	// 利用 type 来判断 field 在不在
	if _, found := typ.FieldByName(field); !found {
		return errors.New("字段不存在")
	}

	if !fd.CanSet() {
		return errors.New("不可修改字段")
	}
	fd.Set(reflect.ValueOf(newVal))
	return nil
}
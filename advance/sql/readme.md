

# SQL 编程


## 初始化 DB

- 驱动和 dsn
		
- "为了这盘醋，包的这顿饺子" 
	    database/sql 定义接口，driver 则是实现接口。通过导包初始化，把驱动实例塞过去


## 增删改查入门

增改删
- Exec 或者 ExecContext，后者可以用来控制超时
- 同时检查 error 和 sql.Result


查询
- QueryRow 和 QueryRowContext，查询单行数据
- Query 和 QueryContext，查询多行数据


要注意参数传递，一般的 SQL 都是使用 `?` 作为参数占位符。
不要把参数拼接进 SQL 本身，容易引起注入。


## Row 和 Rows

Rows
- Rows 迭代器设计，需要在使用前调用 Next 方法
- Scan 支持的类型很多

Row
- 可以理解为只有一行的 Rows，而且是必须要有一行。没有的话，在调用 Row 的 Scan 的时候会返回 sql.ErrNoRow

（批量查询可以没有，单个查询必须要有，想起 gorm 的 Find ）



## driver.Valuer 和 sql.Scanner 接口

场景：SQL 默认支持的类型就是基础类型 []byte 和 string，
该如何自定义类型？比如说我需要支持 json 类型，该如何处理？

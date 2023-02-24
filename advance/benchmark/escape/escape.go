package escape

type User struct {
	Name string
}

func NewUser() *User {
	return &User{
		Name: "Jie",
	}
}



// 闭包引用外部变量
func CountFn() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

package types


type UserInfo struct {
	Name     string  `json:"name"`
	Age      string  `json:"age"`
	IsLeader bool    `json:"is_leader"`
	salary   float64  // 肯定不公开
	// 因为同属一个包，所有 salary 还可以被访问到
	// 如果不同包，就访问不到了
}

func (u *UserInfo) TableName() string {
	return "user"
}

// 访问字段 纠结实例是 UserInfo{} 还是 &UserInfo{}
// 访问方法 略
// 篡改代码 否




// 应用

// 1. 生成代理 - RPC 框架核心 
type UserService struct {
	GetByIdV1 func()  // 可以赋予新的值
}

func (u *UserService) GetByIdV2() {
	// 没办法篡改这个方法
}

func Proxy() {
	myService := &UserService{}
	myService.GetByIdV1 = func() {
		// 发起 RPC 调用
		// 解析响应
	}
}


// 2. 解析模型数据 - Gorm 



// 3. DTO -> PO

// 数据传输
type UserDTO struct {

}

// 数据库直接对应
type UserPO struct {

}

// ignoreFields 忽略一些字段，不拷贝
func Copy(src , dst interface{}, ignoreFields ...string) error {
	// 反射操作，一个个字段拷贝过去
	return nil 
}

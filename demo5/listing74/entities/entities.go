//与人有关类型
package entities

//用户类型
type user struct {
	Name  string
	Email string
}

//管理员
type Admin struct {
	user
	Rights int
}

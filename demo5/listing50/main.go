//将一个类型嵌入另一个类型
package main

import (
	"fmt"
)

//定义一个user用户类型
type user struct {
	name  string
	email string
}

//admin管理员
type admin struct {
	user         //嵌入类型
	level string //级别
}

//实现notify方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {

	//创建一个管理员
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	//可以直接访问内部类型
	ad.user.notify()
	//内部方法可以提升到外部
	ad.notify()
}

/*
Sending user email to john smith<john@yahoo.com>
Sending user email to john smith<john@yahoo.com>
*/

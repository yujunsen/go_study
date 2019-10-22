//内部和外部需要同一个借口时
package main

import (
	"fmt"
)

//定义一个接口类型
type notifier interface {
	notify()
}

//用户类型
type user struct {
	name  string
	email string
}

//调用方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

//管理员
type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
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
	// 给 admin 用户发送一个通知
	// 接口的嵌入的内部类型实现并没有提升到外部内形
	sendNotification(&ad)
	// 我们可以直接访问内部类型的方法
	ad.user.notify()
	// 内部类型的方法没有被提升
	ad.notify()

}

// sendNotification 接受一个实现了 notifier 接口的值
func sendNotification(n notifier) {
	n.notify()
}

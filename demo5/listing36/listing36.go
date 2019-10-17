// Sample program to show how to use an interface in Go.
// 这个示例程序展示 Go 语言里如何使用接口
package main

import "fmt"

// notifier is an interface that defined notification
// type behavior.
// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// user defines a user in the program.
// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
// notify 是使用指针接收者实现的方法
func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	// Create a value of type User and send a notification.
	// notify 是使用指针接收者实现的方法
	u := user{"Bill", "bill@email.com"}
	sendNotification(u)

	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
	// ./listing36.go:32: 不能将 u（类型是 user）作为
	// sendNotification 的参数类型 notifier：
	// user 类型并没有实现 notifier
	// （notify 方法使用指针接收者声明）
}

// sendNotification 接受一个实现了 notifier 接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}

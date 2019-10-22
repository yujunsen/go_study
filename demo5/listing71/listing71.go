package main

import (
	"fmt"

	"github.com/goinaction/mycode/demo5/listing71/entities"
)

func main() {
	// 创建 entities 包中的 User 类型的值
	// u := entities.User{
	// 	Name : "Bill",
	// 	email : "bill@email.com"
	// }
	// 需要把email 改为 Email
	u := entities.User{
		Name:  "Bill",
		Email: "bill@email.com",
	}
	fmt.Printf("User: %v\n", u)
}

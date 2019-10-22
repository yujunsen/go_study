package main

import (
	"fmt"

	"github.com/goinaction/mycode/demo5/listing74/entities"
)

func main() {
	// 创建 entities 包中的 Admin 类型的值
	a := entities.Admin{
		Rights: 10,
	}
	// 设置未公开的内部类型的
	// 公开字段的值
	a.Name = "Bill"
	a.Email = "bill@game.com"
	fmt.Printf("User : %v", a)
}

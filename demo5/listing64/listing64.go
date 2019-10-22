// 这个示例程序展示无法从另一个包里
// 访问未公开的标识符

package main

import (
	"fmt"

	"github.com/goinaction/mycode/demo5/listing64/counters"
)

func main() {
	//创建一个未公开的类型
	//counter := counters.alertCounter(10)
	//.\listing64.go:12:13: undefined: counters 未定义
	conter := counters.GloCounter(10)
	fmt.Printf("conter:%d\n", conter)
}

//当一个标识符的名字以小写字母开头时，这个标识符就是未公开的，即包外的代码不可见。
//如果一个标识符以大写字母开头，这个标识符就是公开的，即被包外的代码可见

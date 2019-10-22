//展示如何使用最基本的 log 包
package main

import "log"

func init() {
	//设置了一个字符串，作为每个日志项的前缀
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")
	// Fatalln 在调用 Println()之后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")
	// Panicln 在调用 Println()之后会接着调用 panic()
	log.Panicln("panic message")
}

/*
const (
// 将下面的位使用或运算符连接在一起，可以控制要输出的信息。没有
// 办法控制这些信息出现的顺序（下面会给出顺序）或者打印的格式
// （格式在注释里描述）。这些项后面会有一个冒号：
// 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
// 日期: 2009/01/23
Ldate = 1 << iota
// 时间: 01:23:23
Ltime
// 毫秒级时间: 01:23:23.123123。该设置会覆盖 Ltime 标志
Lmicroseconds
// 完整路径的文件名和行号: /a/b/c/d.go:23
Llongfile
// 最终的文件名元素和行号: d.go:23
// 覆盖 Llongfile
Lshortfile
// 标准日志记录器的初始值
LstdFlags = Ldate | Ltime
)

关键字 iota 在常量声明区里有特殊的作用。这个关键字让编译器为每个常量复制相同的表
达式，直到声明区结束，或者遇到一个新的赋值语句。关键字 iota 的另一个功能是， iota 的
初始值为 0，之后 iota 的值在每次处理为常量后，都会自增 1

const (
Ldate = 1 << iota // 1 << 0 = 000000001 = 1
Ltime // 1 << 1 = 000000010 = 2
Lmicroseconds // 1 << 2 = 000000100 = 4
Llongfile // 1 << 3 = 000001000 = 8
Lshortfile // 1 << 4 = 000010000 = 16
...
)
*/

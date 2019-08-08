package main

import (
	"log"
	"os"

	_ "github.com/goinaction/mycode/demo2/matchers"
	"github.com/goinaction/mycode/demo2/search"
)

func init() { //在main前先执行
	//log输出为标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("das")
}

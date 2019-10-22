package main

import (
	"fmt"

	"github.com/goinaction/mycode/demo5/listing68/counters"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("counter = %d\n", counter)
}

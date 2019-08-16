package main

import (
	"encoding/json"
	"fmt"
	"os"
	"unsafe"
)

const filename = "test.json"

func writejson() {
	b := []byte(`{"Name":"Wednesday", "Age":6, "Parents": [ "Gomez", "Moticia" ]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		return
	}
	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array", vv)
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}

	}

}

type Text struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Sex   string `json:"sex"`
}

func aa() {

	// file, err := os.Open(filename)
	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open error " + *(*string)(unsafe.Pointer(&err)))
		return
	}
	defer file.Close()

	s := Text{
		Name:  "s",
		Phone: "d",
		Sex:   "x",
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
	println(string(b))
	file.Write(b)
}

func main() {
	//writejson()
	aa()
}

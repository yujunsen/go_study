package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init is called before main.
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:./exmaple2 <url>")
		os.Exit(-1)
	}
}

// main is the entry point for the application.
func main() {
	// Get a response from the web server.
	r, err := http.Get(os.Args[1])
	if err != nil {
		//fmt.Println(os.Args[1])
		fmt.Println(err)
		return
	}

	// Copies from the Body to Stdout.
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

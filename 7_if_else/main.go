package main

import "fmt"

func main() {

	age := 10

	if age > 18 {
		fmt.Println("Hello, Buddy You are an Adult")
	} else if age > 12 {
		fmt.Println("hello, buddy you are a teenager")
	} else {
		fmt.Println("Hello buddy you are a kid")
	}
}

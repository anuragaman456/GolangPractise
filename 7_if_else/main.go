package main

import "fmt"

func main() {

	age := 20

	if age > 18 {
		fmt.Println("Hello, Buddy You are an Adult")
	} else if age > 12 {
		fmt.Println("hello, buddy you are a teenager")
	} else {
		fmt.Println("Hello buddy you are a kid")
	}

	// include logical operator inside if else condition

	Role := "admin"
	Access := true

	if Role == "admin" && Access == true {
		fmt.Println("Access Allowed")
	} else {
		fmt.Println("Access Denied!")
	}

	// you can declare variable inside if also

	if age := 13; age >= 18 {
		fmt.Println("Hello There You are adult, your age is:", age)
	} else if age > 10 {
		fmt.Println("Hello there you are Teenager, your age is:", age)
	} else {
		fmt.Println("Hello There you are Kid, your age is:", age)
	}

	// there is no concept of ternary operaotor in Golang

}

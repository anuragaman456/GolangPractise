package main

import "fmt"

var name = "Anurag Aman"

func main() {
	var num int = 6
	var num2 = 9 // agr ap type define nahi kroge to go khud type define kr dega

	var num3 = "9" // yha maine type nahi diya wo automaticaly string define kr dega

	var num4 float64 = 10.8

	num5 := 10.99 // shortest way to initialise and create value

	// But we cant do this short form and without type when we have to initialise variable value later

	var num6 string

	num6 = "Anurag Aman"

	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(name)

}

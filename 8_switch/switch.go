package main

import "fmt"

func main() {
	i := 9
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("five")

	default:
		fmt.Println("pLease Enter under Range!")
	}
}

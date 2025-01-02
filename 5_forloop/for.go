package main

import "fmt"

func main() {

	var i int = 0
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	//create infinite loop

	// for {
	// 	fmt.Println(1)
	// }

	// classic for loop used in programming

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

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

	// added continue( to skip some value or iteration ) in loop
	// added break to end the loop
	for i := 100; i <= 110; i++ {
		if i == 101 {
			continue
		}
		if i == 108 {
			fmt.Println(i)
			break
		}
		fmt.Println(i)

	}
}

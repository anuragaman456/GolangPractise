package main

import "fmt"

func main() {
	const name string = "Shivam Bhardwaj"

	const ( // we acn write multipe const together also
		version="2.0"
		port = 5000
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(port)
	fmt.Println(port, host)

}

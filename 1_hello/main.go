package main // basically it tell that it is entry of the go file and without this u cant able to make executable file

import "fmt" //The import statement brings in external libraries or packages into your code. Here, "fmt" isa standard Go library used for formatted I/O (like printing to the console)

func main() { // The Go runtime requires a main function in the main package to begin execution. It’s where the program logic is implemented.
	fmt.Println("Hello World3") // from external fmt package it is importing println to console the output in cli

}

// two ways of execution -{1)go build main.go then ./main.exe   2) use directly go run main.go}

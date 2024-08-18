package main

import "fmt"

func main() {
	for {
		var input string
		fmt.Scanln(&input)
		fmt.Printf("string: %s\n", input)
	}
}

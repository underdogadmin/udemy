package main

import "fmt"

func main() {

	var small int
	var big int

	fmt.Print("Enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Enter a big number: ")
	fmt.Scan(&big)

	fmt.Println(big%small, "is the remainder of", big, "divided by", small)
}

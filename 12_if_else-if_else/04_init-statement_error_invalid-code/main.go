package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// 	fmt.Println(food)     // would be invalid to put here since it's not inside the correct scope
}

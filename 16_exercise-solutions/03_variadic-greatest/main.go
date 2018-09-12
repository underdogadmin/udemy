package main

import "fmt"

func main() {
	n := max(1, 2, 3, 4, 5)
	fmt.Println(n)
}

func max(z ...int) int {
	var biggest int
	for _, r := range z {
		if r > biggest {
			biggest = r
		}
	}
	return biggest
}

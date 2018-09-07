package main

import "fmt"

func main() {

	myFreindsName := "Mar"

	switch {
	case len(myFreindsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFreindsName == "Tim":
		fmt.Println("Wassup Time")
	case myFreindsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFreindsName == "Marcus", myFreindsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFreindsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFreindsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}

/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions
*/

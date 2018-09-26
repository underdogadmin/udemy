package main

import (
	"encoding/json"
	"fmt"
)

// Person - adding a comment so VS Code doesnt bark at me
type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}

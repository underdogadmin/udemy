package main

import (
	"encoding/json"
	"os"
)

// Person - adding a comment so VS Code doesnt bark at me
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}

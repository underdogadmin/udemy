package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

// HashBucket function comment so VS code doesnt yell at me
func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

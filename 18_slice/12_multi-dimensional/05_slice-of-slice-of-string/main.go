package main

import "fmt"

func main() {
	records := make([][]string, 0)
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	studen2 := make([]string, 4)
	studen2[0] = "Gomez"
	studen2[1] = "Lisa"
	studen2[2] = "92.00"
	studen2[3] = "96.00"
	// store the record
	records = append(records, studen2)
	// print
	fmt.Println(records)
}

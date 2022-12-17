package main

import "fmt"

// Function to print the state of the statement
// Prints the state of the environment S
func printValueStatemnt(s ValState) {
	fmt.Printf("\n S: [\t") // Begin the environment S
	for k, v := range s {
		if v.flag == ValueInt {
			fmt.Printf(k, " : ", v.valI, "\t")
		} else if v.flag == ValueBool {
			fmt.Printf(k, " : ", v.valB, "\t")
		} else {
			fmt.Printf("Undefined value in S: %s", k)
		}
	}
	fmt.Printf("]\t") // End the environment S
}

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

// Same as printValueStatemnt but for types
func printTypeStatemnt(t TyState) bool {
	fmt.Printf("\n G: [\t") // Begin the environment T
	for k, v := range t {
		if v == TyInt {
			fmt.Printf(k, " : ", "int", "\t")
		}
		if v == TyBool {
			fmt.Printf(k, " : ", "bool", "\t")
		} else {
			fmt.Printf("Undefined type in G: %s", k)
			return false
		}
	}
	fmt.Printf("]\t") // End the environment T
	return true
}

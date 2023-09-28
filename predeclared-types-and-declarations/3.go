// Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64. Assign each variable the maximum legal value for its type, then add 1 to each variable. Print out their values.

package main

import "fmt"

func check_type_limits() (byte, int32, uint64) {
	// Assigning all the variables with highest value of their corresponding types.
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println("value of b =", b)
	fmt.Println("value of smallI =", smallI)
	fmt.Println("value of bigI =", bigI)

	// Incrementing all the variables.
	b++
	smallI++
	bigI++
	fmt.Println("")
	fmt.Println("value of b after increment =", b)
	fmt.Println("value of smallI after increment =", smallI)
	fmt.Println("value of bigI after increment =", bigI)

	return b, smallI, bigI
}

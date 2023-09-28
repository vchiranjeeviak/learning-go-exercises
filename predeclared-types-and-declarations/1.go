// Write a program where you declare an integer variable called i with the value 20. Assign i to a floating point variable named f. Print out i and f.
package main

import "fmt"

func assign_int_to_float() (int, float64){
	var i int = 20
	// Explicit type conversion. Only supported types can be converted like this. String 
	var f float64 = float64(i)

	fmt.Println("value of i =",i)
	fmt.Println("value of f =",f)

	return i, f
}
// Write a program where you declare a constant called value that can be assigned to both an integer and a floating point variable. Assign it to an integer called i and a floating point variable called f. Print out i and f.

package main

import "fmt"

func assign_constant_to_int_and_float() (int, float64) {
	// A literal doesn't have a type if not mentioned while declaring it until it is assigned to something or used in some context.
	// Here 30 has no type.
	const value = 30

	// Here 30 is of type int
	var i int = value
	// Here 30 is of type float64
  var f float64 = value

	fmt.Println("value of i =",i)
	fmt.Println("value of f =",f)
	
	return i, f
}
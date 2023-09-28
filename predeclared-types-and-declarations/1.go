// Write a program where you declare an integer variable called i with the value 20. Assign i to a floating point variable named f. Print out i and f.
package types

import "fmt"

func assign_int_to_float() (int, float64){
	var i int = 20
	var f float64 = float64(i)

	fmt.Println("value of i =",i)
	fmt.Println("value of f =",f)

	return i, f
}
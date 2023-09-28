// Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". Create a sub-slice containing the first two values, a second subslice with the second, third, and fourth values, and a third subslice with the fourth and fifth values. Print out all four slices.

package main

import "fmt"

type stringSlice []string

func create_slice_and_subslices() (stringSlice, stringSlice, stringSlice, stringSlice) {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	subslice1 := greetings[:2]
	subslice2 := greetings[1:4]
	subslice3 := greetings[3:]

	fmt.Println("Greetings =", greetings)
	fmt.Println("Subslice1 =", subslice1)
	fmt.Println("Subslice2 =", subslice2)
	fmt.Println("Subslice3 =", subslice3)

	return greetings, subslice1, subslice2, subslice3
}

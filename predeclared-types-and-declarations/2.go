// Write a program where you declare a constant called value that can be assigned to both an integer and a floating point variable. Assign it to an integer called i and a floating point variable called f. Print out i and f.

package types

func assign_constant_to_int_and_float() (int, float64) {
	const value = 30

	var i int = value
  var f float64 = value

	return i, f
}
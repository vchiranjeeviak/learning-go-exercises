package main

import (
	"reflect"
	"testing"
)

func Test_assign_int_to_float(t *testing.T) {
	i, f := assign_int_to_float()

	type_of_i := reflect.TypeOf(i).String()
	if type_of_i != "int" {
		t.Errorf("Expected type int for variable i, but got %s instead", type_of_i)
	}

	type_of_f := reflect.TypeOf(f).String()
	if type_of_f != "float64" {
		t.Errorf("Expected type float64 for variable f, but got %s instead", type_of_f)
	}
}

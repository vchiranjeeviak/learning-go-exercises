package main

import "testing"

func Test_check_type_limits(t *testing.T) {
	b, smallI, bigI := check_type_limits()

	if b != 0 {
		t.Error("Byte b didn't wrapped around, check the value.")
	}

	if smallI != -2147483648 {
		t.Error("Int32 smallI didn't wrap around, check the value.")
	}

	if bigI != 0 {
		t.Error("Uint64 bigI didn't wrap around, check the value.")
	}
}

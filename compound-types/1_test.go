package main

import "testing"

func compareStringSlices(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	for i, v := range first {
		if v != second[i] {
			return false
		}
	}

	return true
}

func Test_create_slice_and_subslices(t *testing.T) {
	greetings_test := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subslice1_test := []string{"Hello", "Hola"}
	subslice2_test := []string{"Hola", "नमस्कार", "こんにちは"}
	subslice3_test := []string{"こんにちは", "Привіт"}
	greetings, subslice1, subslice2, subslice3 := create_slice_and_subslices()

	if !compareStringSlices(greetings, greetings_test) {
		t.Error("Source slice is not correct. Check in the actual file.")
	}

	if !compareStringSlices(subslice1, subslice1_test) {
		t.Error("Subslice1 is incorrect, check the logic.")
	}

	if !compareStringSlices(subslice2, subslice2_test) {
		t.Error("Subslice2 is incorrect, check the logic.")
	}

	if !compareStringSlices(subslice3, subslice3_test) {
		t.Error("Subslice3 is incorrect, check the logic.")
	}
}

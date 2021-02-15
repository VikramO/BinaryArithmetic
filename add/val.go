package main

import (
	"fmt"
	"strings"
)

func Validate(num1 string, num2 string) bool {

	valid_digits := "01"

	var valid bool

	len_1 := len([]rune(num1))
	len_2 := len([]rune(num2))

	max_length := Max(len_1, len_2)
	min_length := Min(len_1, len_2)

	diff := max_length - min_length

	if len_1 > len_2 {
		append := strings.Repeat("0", diff)
		num2 = num2 + append
	} else if len_1 < len_2 {
		append := strings.Repeat("0", diff)
		num1 = num1 + append
	}

	for i := 0; i < max_length; i++ {
		found_1 := strings.Index(valid_digits, string([]rune(num1)[i]))
		found_2 := strings.Index(valid_digits, string([]rune(num2)[i]))

		if found_1 < 0 {
			fmt.Println("Your first input is invalid")
			valid = false
			return valid
		} else if found_2 < 0 {
			fmt.Println("Your second input is invalid")
			valid = false
			return valid
		}
	}

	fmt.Println("Your inputs are valid, computing sum")
	valid = true
	return valid
}

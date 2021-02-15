package main

import (
	"fmt"
)

func main() {
	num1 := "101011010"
	num2 := "10011101"
	valid := Validate(num1, num2)
	if valid == true {
		sum := Add(num1, num2)
		fmt.Println("The sum in binary equals" + " " + sum)
	} else {
	}
}

//Output:
//Your inputs are valid, computing sum
//The sum in binary equals 111110111

//If either num1 or num2 contains a digit other than 0 or 1:
//Output:
//Your {first/second} input is invalid

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
		print(valid)
	}
}

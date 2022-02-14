// function with multiple return values
package main

import "fmt"

func Calculator(a, b int) (int, int, int, int, int) {
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	moddiv := a % b
	return sum, sub, mul, div, moddiv
}

func main() {
	sum, sub, mul, div, moddiv := Calculator(34, 23)
	fmt.Println("the sum of 34 and 23 is", sum, "\nthe subraction of 34 and 23 is", sub, "\nthe multiplication of 34 and 23 is", mul, "\nthe division of 34 and 23 is", div, "\nthe modulodivision of 34 and 23 is", moddiv)
}

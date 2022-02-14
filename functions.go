// normal function declaration program

package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func modiv(a, b int) int {
	return a % b
}

func main() {
	op1 := sum(15, 34)
	op2 := sub(15, 34)
	op3 := mul(15, 34)
	op4 := modiv(15, 34)
	fmt.Println("the sum of two numbers are", op1)
	fmt.Println("the sub of two numbers are", op2)
	fmt.Println("the mul of two numbers are", op3)
	fmt.Println("the modulo division of two numbers are", op4)
}

//----------------------------------------------------------------------------------------------------------

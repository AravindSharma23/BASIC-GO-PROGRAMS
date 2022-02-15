package main

import "fmt"

func main() {
	s := []int{23, 12, 34, 45, 66, 77, 88, 98} // slice
	defer fmt.Println(s)
	frds := make([]string, 5, 10)
     frds[0] = "aravind"
	 frds[1] = "ravi"
	 frds[2] = "kumar"
	 frds[3]= "gokul"
	 frds[4] = "sakthi"
	fmt.Println(frds)
	defer foo()
	bar()

}
func foo() {
	fmt.Println("Hello from foo function")
}
func bar() {
	fmt.Println("Greetings from bar function")
}

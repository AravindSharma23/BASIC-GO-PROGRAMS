package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{
		Name: "Aravind",
		Age:  21,
	}
	u2 := User{
		Name: "kumar",
		Age:  31,
	}
	u3 := User{
		Name: "Ramesh",
		Age:  54,
	}
	u4 := User{
		Name: "Karthi",
		Age:  37,
	}
	Consumers := []User{u1, u2, u3, u4}
	fmt.Println(Consumers)
	e, err := json.Marshal(Consumers)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(string(e))
}

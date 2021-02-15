package main

import (
	"fmt"

	"example.com/mathpack"
)

func main() {
	fmt.Println("Start")
	u := User{name: "Test User"}
	fmt.Println(u)
	i := 5.0
	m := 2.0
	fmt.Println(mathpack.Pow(i, m))
}

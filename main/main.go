package main

import (
	"fmt"

	"github.com/cholladay0816/gotest/mathpack"
	"github.com/cholladay0816/gotest/models"
)

func main() {
	fmt.Println("Start")
	u := models.User{Name: "Test User"}
	fmt.Println(u)
	i := 5.0
	m := 2.0
	fmt.Println(mathpack.Pow(i, m))
}

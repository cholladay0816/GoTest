package main

import (
	"fmt"
)

var arr [10]int

func main() {
	for i := 0; i < 10; i++ {
		arr[i] = (i + 1) * 2
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

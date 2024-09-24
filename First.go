package main

import (
	"fmt"
)

func main() {
	var Name string
	fmt.Println("Please Enter your Name: ")
	fmt.Scanln(&Name)

	fmt.Printf("Hello %s", Name)
}

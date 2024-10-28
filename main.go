package main

import (
	"fmt"
)

func main() {
	var Name string
	fmt.Println("Please Enter your Name: ")
	_, err := fmt.Scanln(&Name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Hello %s", Name)
}

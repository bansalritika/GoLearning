package main

import "fmt"

func main() {
	mynumber := 23
	ptr := &mynumber // refernce the value
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr + 2
	fmt.Println(*ptr)
}
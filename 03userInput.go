package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter Your Name")
	reader := bufio.NewReader(os.Stdin)
	
	//comma ok syntax
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello,", strings.TrimSpace(name))
}
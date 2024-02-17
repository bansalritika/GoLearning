package main

import "fmt"

//constant : if constname is with capital letter than it will be public
const Login string = "ritika is login"

func main() {
	//string
	var name string = "Ritika"
	fmt.Println(name)
	fmt.Printf("the data type of name is: %T \n", name)

	//boolean
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("the data type is: %T\n", isLoggedIn)

	//uint for unsigned integers
	var smallvalue uint8 = 254
	fmt.Println(smallvalue)
	fmt.Printf("the data type is : %T \n", smallvalue)

	//float
	var floatvalue float32 = 255.768234
	fmt.Println(floatvalue)
	fmt.Printf("the data type is : %T \n", floatvalue)

	//default values and somealiases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("the data type is : %T \n", anothervariable) //value will be 0

	var stringvariable string
	fmt.Println(stringvariable)
	fmt.Printf("the data type is: %T \n", stringvariable)

	//implicit type
	var website = "www.google.com"
	fmt.Println(website)

	//no var These walrus operator are not allowed outside a method
	age := 18
	fmt.Println(age)

	//const
	fmt.Println(Login)
}

package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Apple"
	fruitList[2] = "Apple"
	fmt.Println("Fruitlist is :", fruitList)
	fmt.Println("length is", len(fruitList))


	var vegList = [4]string {"Potato", "Brocolli", "Beans"}
	fmt.Println("VegList is: ", vegList)


	myArr := [...]int{10, 20, 30, 40, 50}
	for i := 0; i < len(myArr); i++ {
		fmt.Print(myArr[i], " ")
	}
}
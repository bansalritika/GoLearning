package main

import (
	"fmt"
	"sort"
)

func main() {
	var mobilePhones = []string{"Apple", "RealMe", "Nokia"}
	fmt.Printf("Type of mobilePhones is: %T\n", mobilePhones)

	mobilePhones = append(mobilePhones, "Oppo", "Poco")
	fmt.Println(mobilePhones)

	mobilePhones = append(mobilePhones[0:2])
	fmt.Println(mobilePhones)

	highScores := make([]int , 3)
	highScores[0] = 234
	highScores[1] = 334
	highScores[2] = 434
	highScores = append(highScores, 534, 634)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//REMOVE A VALUE FROM SLICE BASED ON INDEX

	courses := []string{"python", "javascript", "ruby", "swift", "nodejs"}
	fmt.Println(courses)
	// index := 2
	courses = append(courses[:2], courses[2 + 1:]...)
	fmt.Println(courses)
}
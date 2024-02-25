package main

import "fmt"

func main() {
	//1
	vegList := map[string]int{
		"Brocolli": 10,
		"Beans":    100,
		"Potato":   5,
	}

	for key, value := range vegList {
		fmt.Println(key, " ", value)
	}

	//2
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"
	fmt.Println(languages)

	delete(languages, "JS")
	fmt.Println(languages)

	//loop
	for key, value := range languages {
		fmt.Printf("for key %v , value is %v \n", key, value)
	}

	//found
	_, found := languages["JS"]
	if !found {
		fmt.Println("language not found")
	} else {
			fmt.Println("language found")
		}
	
}

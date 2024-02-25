package main

import (
	"fmt"
	"time"
)

func listener(data chan <- string) {
	for {
		fmt.Println("Received ", <-data)
	}
}

func sender(data chan <- string) {
	
}

func main() {
	data := make(chan string)
	data <- "Hello, world!"
	go listener(data)
	go sender(data)

	// Sleep for a short period of time to allow the goroutines to run
	time.Sleep(1 * time.Second)
} 
// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	result := make([]int, len(nums))
// 	var wg sync.WaitGroup
// 	for _, num := range nums {
// 		wg.Add(1)
// //run 2 routines in one set
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to bubble sort")
	// get the random slice to process for the bubble sort
	unsortedSlice := generateRandomSlice(50)
	fmt.Println("Unsorted Randome array before the sortirng", unsortedSlice)
	// call bubble sorting fuction with unsorted slice
	bubbleSorter(&unsortedSlice)
	// Print the sorted array
	fmt.Println("Sorted Array is ", unsortedSlice)
}

func generateRandomSlice(sizeOfSlice int) []int {
	//create a slice with given size
	randomSlice := make([]int, sizeOfSlice, sizeOfSlice)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sizeOfSlice; i++ {
		randomSlice[i] = rand.Intn(99)
	}
	return randomSlice
}

func bubbleSorter(unSorted *[]int) {
	lenOfSlice := len(*unSorted)
	breaker := lenOfSlice - 1
	for {
		if breaker == 0 {
			break
		}
		for i := 0; i < lenOfSlice-1; i++ {
			if (*unSorted)[i] > (*unSorted)[i+1] {
				(*unSorted)[i], (*unSorted)[i+1] = (*unSorted)[i+1], (*unSorted)[i]
			}
		}
		breaker -= 1
	}

}
